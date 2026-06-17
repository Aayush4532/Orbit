package buyer

import "time"

type PurchaseResponse struct {
	ReservationID string     `json:"reservationId"`
	Price         float64    `json:"price"`
	Status        string     `json:"status"`
	ExpiresAt     *time.Time `json:"expiresAt,omitempty"`
}