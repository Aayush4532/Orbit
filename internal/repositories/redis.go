package repositories

import (
	"Orbit/internal/db"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type BookingResult int

const (
	BookingSoldOut        BookingResult = 0
	BookingSuccess        BookingResult = 1
	BookingAlreadyDone    BookingResult = 2
	BookingProductMissing BookingResult = 3
)

var ErrProductNotFound = errors.New("product not found in inventory")

type Reservation struct {
	ReservationID string
	ProductID     string
	EventID       string
	UserID        string
	Price         float64
	Status        string
	BookedAt      time.Time
	ExpiresAt     time.Time // zero value = no expiry
}

var reserveScript = redis.NewScript(`
	-- Idempotency: booking hash existence = already booked
	if redis.call('EXISTS', KEYS[2]) == 1 then
		return 2
	end

	-- Stock + price from product hash (single round trip)
	local fields = redis.call('HMGET', KEYS[1], 'stock', 'price')
	local stock = tonumber(fields[1])
	local price = fields[2]

	-- Product not loaded into Redis at all
	if not stock and not price then
		return 3
	end

	-- Product exists but stock is zero
	if not stock or stock <= 0 then
		return 0
	end

	-- Claim one unit
	redis.call('HINCRBY', KEYS[1], 'stock', -1)

	-- Booking record: full audit trail + idempotency in one key
	redis.call('HSET', KEYS[2],
		'userId',    ARGV[1],
		'productId', ARGV[2],
		'eventId',   ARGV[3],
		'price',     price,
		'status',    ARGV[6],
		'bookedAt',  ARGV[4]
	)

	-- TTL only when payment window is set.
	-- CRITICAL: EXPIRE key 0 DELETES the key — must be conditional.
	local ttl = tonumber(ARGV[5])
	if ttl > 0 then
		redis.call('EXPIRE', KEYS[2], ttl)
	end

	-- Audit trail (permanent — no TTL on this set)
	redis.call('SADD', KEYS[3], ARGV[1])

	return 1
`)

func ReserveProduct(
	ctx context.Context,
	productId string,
	eventId string,
	userId string,
	status string,
	ttl time.Duration,
) (*Reservation, BookingResult, error) {
	rdb := db.GetRedisClient()

	productKey := fmt.Sprintf("product:%s:%s", productId, eventId)
	bookingKey := fmt.Sprintf("booking:%s:%s:%s", userId, productId, eventId)
	winnersKey := fmt.Sprintf("event:%s:product:%s:winners", eventId, productId)

	now := time.Now()
	ttlSeconds := fmt.Sprintf("%d", int64(ttl.Seconds()))

	keys := []string{productKey, bookingKey, winnersKey}
	args := []interface{}{
		userId,
		productId,
		eventId,
		fmt.Sprintf("%d", now.Unix()),
		ttlSeconds,
		status,
	}

	raw, err := reserveScript.Run(ctx, rdb, keys, args...).Int()
	if err != nil {
		return nil, BookingSoldOut, fmt.Errorf("reserve script: %w", err)
	}

	result := BookingResult(raw)
	if result != BookingSuccess {
		return nil, result, nil
	}
	priceStr, err := rdb.HGet(ctx, bookingKey, "price").Result()
	if err != nil {
		return nil, BookingSuccess, fmt.Errorf("fetch booking price: %w", err)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, BookingSuccess, fmt.Errorf("parse price %q: %w", priceStr, err)
	}

	reservation := &Reservation{
		ReservationID: bookingKey, 
		ProductID:     productId,
		EventID:       eventId,
		UserID:        userId,
		Price:         price,
		Status:        status,
		BookedAt:      now,
	}

	if ttl > 0 {
		reservation.ExpiresAt = now.Add(ttl)
	}

	return reservation, BookingSuccess, nil
}
