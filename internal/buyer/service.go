package buyer

import (
	"Orbit/internal/repositories"
	"Orbit/internal/utils"
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrSoldOut         = errors.New("sold out")
	ErrAlreadyBooked   = errors.New("already booked")
	ErrProductNotFound = errors.New("product not found")
)

type Service struct {
	requirePayment bool
	paymentWindow  time.Duration
}

func NewService(requirePayment bool, paymentWindow time.Duration) *Service {
	return &Service{
		requirePayment: requirePayment,
		paymentWindow:  paymentWindow,
	}
}

func (s *Service) Buy(ctx context.Context, productId, eventId string, claim *utils.Claims) (*PurchaseResponse, error) {
	status := StatusConfirmed
	ttl := time.Duration(0)

	if s.requirePayment {
		status = StatusPendingPayment
		ttl = s.paymentWindow
	}

	reservation, result, err := repositories.ReserveProduct(
		ctx,
		productId,
		eventId,
		claim.ID,
		string(status),
		ttl,
	)
	if err != nil {
		return nil, fmt.Errorf("buyer service: %w", err)
	}

	switch result {
	case repositories.BookingSoldOut:
		return nil, ErrSoldOut
	case repositories.BookingAlreadyDone:
		return nil, ErrAlreadyBooked
	case repositories.BookingProductMissing:
		return nil, ErrProductNotFound
	case repositories.BookingSuccess:
		resp := &PurchaseResponse{
			ReservationID: reservation.ReservationID,
			Price:         reservation.Price,
			Status:        reservation.Status,
		}
		if !reservation.ExpiresAt.IsZero() {
			resp.ExpiresAt = &reservation.ExpiresAt
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("buyer service: unexpected result %d", result)
	}
}