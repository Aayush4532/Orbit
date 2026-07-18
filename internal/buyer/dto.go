package buyer

import "time"

type BookingStatus string

const (
	StatusConfirmed      BookingStatus = "CONFIRMED"
	StatusPendingPayment BookingStatus = "PENDING_PAYMENT"
)

type PurchaseResponse struct {
	ReservationID string     `json:"reservationId"`
	Price         float64    `json:"price"`
	Status        string     `json:"status"`
	ExpiresAt     *time.Time `json:"expiresAt,omitempty"` // only set for PENDING_PAYMENT
}