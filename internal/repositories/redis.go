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

var (
	ErrProductNotFound = errors.New("product not found in inventory")
	ErrSoldOut         = errors.New("sold out")
	ErrAlreadyBooked   = errors.New("already booked")
)

type BookingResult int

const (
	BookingSoldOut        BookingResult = 0
	BookingSuccess        BookingResult = 1
	BookingAlreadyDone    BookingResult = 2
	BookingProductMissing BookingResult = -3
)

type Reservation struct {
	ReservationID string
	Price         float64
	Status        string
	ExpiresAt     time.Time
}

var reserveScript = redis.NewScript(`
local existing = redis.call('GET', KEYS[4])
if existing then
	return {2, existing, ''}
end

local stock = tonumber(redis.call('GET', KEYS[1]))
if not stock or stock <= 0 then
	return {0, '', ''}
end

local price = redis.call('HGET', KEYS[2], 'price')
if not price then
	return {-3, '', ''}
end

redis.call('DECR', KEYS[1])

local reservationId = redis.call('INCR', 'reservation:seq')
local ttl = tonumber(ARGV[6])

redis.call('HSET', KEYS[3],
	'reservationId', tostring(reservationId),
	'userId', ARGV[1],
	'productId', ARGV[2],
	'eventId', ARGV[3],
	'price', price,
	'status', ARGV[4],
	'createdAt', ARGV[5]
)

if ttl and ttl > 0 then
	redis.call('EXPIRE', KEYS[3], ttl)
	redis.call('SET', KEYS[4], tostring(reservationId), 'EX', ttl)
else
	redis.call('SET', KEYS[4], tostring(reservationId))
end

redis.call('SADD', KEYS[5], ARGV[1])

return {1, tostring(reservationId), price}
`)

func ReserveProduct(
	ctx context.Context,
	productID string,
	eventID string,
	userID string,
	status string,
	ttl time.Duration,
) (Reservation, BookingResult, error) {
	rdb := db.GetRedisClient()

	stockKey := fmt.Sprintf("product:%s:%s", productID, eventID)
	metaKey := fmt.Sprintf("productmeta:%s:%s", productID, eventID)
	reservationKey := fmt.Sprintf("reservation:%s:%s:%s", userID, productID, eventID)
	idempotencyKey := fmt.Sprintf("booked:%s:%s:%s", userID, productID, eventID)
	winnersKey := fmt.Sprintf("event:%s:product:%s:winners", eventID, productID)

	ttlSeconds := int64(ttl / time.Second)
	if ttl > 0 && ttlSeconds == 0 {
		ttlSeconds = 1
	}

	raw, err := reserveScript.Run(
		ctx,
		rdb,
		[]string{stockKey, metaKey, reservationKey, idempotencyKey, winnersKey},
		userID,
		productID,
		eventID,
		status,
		time.Now().UTC().Format(time.RFC3339Nano),
		strconv.FormatInt(ttlSeconds, 10),
	).Result()
	if err != nil {
		return Reservation{}, BookingSoldOut, fmt.Errorf("booking script: %w", err)
	}

	vals, ok := raw.([]interface{})
	if !ok || len(vals) < 3 {
		return Reservation{}, BookingSoldOut, fmt.Errorf("unexpected lua response: %T", raw)
	}

	code, err := toInt64(vals[0])
	if err != nil {
		return Reservation{}, BookingSoldOut, err
	}

	result := BookingResult(code)

	switch result {
	case BookingSoldOut:
		return Reservation{}, result, ErrSoldOut

	case BookingAlreadyDone:
		reservationID, _ := toString(vals[1])
		return Reservation{ReservationID: reservationID, Status: status}, result, ErrAlreadyBooked

	case BookingProductMissing:
		return Reservation{}, result, ErrProductNotFound

	case BookingSuccess:
		reservationID, _ := toString(vals[1])
		priceStr, _ := toString(vals[2])
		price, _ := strconv.ParseFloat(priceStr, 64)

		res := Reservation{
			ReservationID: reservationID,
			Price:         price,
			Status:        status,
		}
		if ttl > 0 {
			res.ExpiresAt = time.Now().UTC().Add(ttl)
		}
		return res, result, nil

	default:
		return Reservation{}, result, fmt.Errorf("unknown booking result: %d", code)
	}
}

func toInt64(v any) (int64, error) {
	switch t := v.(type) {
	case int64:
		return t, nil
	case int:
		return int64(t), nil
	case string:
		return strconv.ParseInt(t, 10, 64)
	case []byte:
		return strconv.ParseInt(string(t), 10, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", v)
	}
}

func toString(v any) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case []byte:
		return string(t), nil
	case int64:
		return strconv.FormatInt(t, 10), nil
	case int:
		return strconv.Itoa(t), nil
	default:
		return "", fmt.Errorf("cannot convert %T to string", v)
	}
}