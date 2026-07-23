package buyer

import (
	"Orbit/internal/models"
	"Orbit/internal/repositories"
	"Orbit/internal/utils"
	"context"
	"errors"
	"fmt"
	// "log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	ErrSoldOut         = errors.New("sold out")
	ErrAlreadyBooked   = errors.New("already booked")
	ErrProductNotFound = errors.New("product not found")
	ErrSaleNotActive   = errors.New("sale is paused or has ended")
	ErrInvalidEventID  = errors.New("invalid event ID")
)

type Service struct {
	requirePayment bool
	paymentWindow  time.Duration
}

func NewService(requirePayment bool, paymentWindow time.Duration) *Service {
	return &Service{requirePayment: requirePayment, paymentWindow: paymentWindow}
}

func (s *Service) Buy(ctx context.Context, productId, eventId string, claim *utils.Claims) (*PurchaseResponse, error) {
	status := StatusConfirmed
	ttl := time.Duration(0)
	if s.requirePayment {
		status = StatusPendingPayment
		ttl = s.paymentWindow
	}

	reservation, result, err := repositories.ReserveProduct(ctx, productId, eventId, claim.ID, string(status), ttl)
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
	case repositories.BookingSaleNotActive:
		return nil, ErrSaleNotActive
	case repositories.BookingSuccess:
		// go s.persistOrder(reservation, productId, eventId, claim.ID)

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

// func (s *Service) persistOrder(reservation *repositories.Reservation, productId, eventId, userId string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	userObjId, err := bson.ObjectIDFromHex(userId)
// 	if err != nil {
// 		log.Printf("WARN persistOrder: invalid userId %s: %v", userId, err)
// 		return
// 	}
// 	productObjId, err := bson.ObjectIDFromHex(productId)
// 	if err != nil {
// 		log.Printf("WARN persistOrder: invalid productId %s: %v", productId, err)
// 		return
// 	}
// 	eventObjId, err := bson.ObjectIDFromHex(eventId)
// 	if err != nil {
// 		log.Printf("WARN persistOrder: invalid eventId %s: %v", eventId, err)
// 		return
// 	}

// 	now := time.Now()
// 	order := models.Order{
// 		ID:            bson.NewObjectID(),
// 		UserID:        userObjId,
// 		ProductID:     productObjId,
// 		EventID:       eventObjId,
// 		ReservationID: reservation.ReservationID,
// 		Price:         reservation.Price,
// 		Status:        models.OrderStatus(reservation.Status),
// 		CreatedAt:     now,
// 		UpdatedAt:     now,
// 	}

// 	if err := repositories.CreateOrder(ctx, order); err != nil {
// 		log.Printf("WARN persistOrder: save failed for reservation=%s: %v", reservation.ReservationID, err)
// 	}
// }

func (s *Service) GetLiveEvents(ctx context.Context) ([]EventView, error) {
	events, err := repositories.GetLiveEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("get live events: %w", err)
	}
	views := make([]EventView, len(events))
	for i, e := range events {
		views[i] = EventView{
			ID:          e.ID.Hex(),
			EventName:   e.EventName,
			Description: e.Description,
			ImageBanner: e.ImageBanner,
			ScheduledAt: e.ScheduledAt,
		}
	}
	return views, nil
}

func (s *Service) GetEventProducts(ctx context.Context, eventIdStr string) ([]ProductView, error) {
	eventId, err := bson.ObjectIDFromHex(eventIdStr)
	if err != nil {
		return nil, ErrInvalidEventID
	}
	products, err := repositories.GetEventProductsWithStock(ctx, eventId)
	if err != nil {
		return nil, fmt.Errorf("get event products: %w", err)
	}
	views := make([]ProductView, len(products))
	for i, p := range products {
		views[i] = ProductView{
			ProductID:      p.ID.Hex(),
			Title:          p.Title,
			Description:    p.Description,
			Price:          p.Price,
			Currency:       p.Currency,
			Image:          p.Image,
			AvailableStock: p.AvailableStock,
		}
	}
	return views, nil
}

func (s *Service) GetMyOrders(ctx context.Context, userId string) ([]models.Order, error) {
	userObjId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	return repositories.GetOrdersByUser(ctx, userObjId)
}
