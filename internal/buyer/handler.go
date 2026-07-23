package buyer

import (
	"Orbit/internal/utils"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) BuyerEventHandler(c *gin.Context) {
	productId := c.Param("productId")
	eventId := c.Param("eventId")

	rawClaim, ok := c.Get("UserFields")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "not permitted"})
		return
	}
	userClaim, ok := rawClaim.(*utils.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid session"})
		return
	}
	if !userClaim.IsApproved {
		c.JSON(http.StatusForbidden, gin.H{"error": "verification required"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 500*time.Millisecond)
	defer cancel()

	resp, err := h.svc.Buy(ctx, productId, eventId, userClaim)
	if err != nil {
		switch {
		case errors.Is(err, ErrSoldOut):
			c.JSON(http.StatusConflict, gin.H{"error": "sold out"})
		case errors.Is(err, ErrAlreadyBooked):
			c.JSON(http.StatusConflict, gin.H{"error": "already booked"})
		case errors.Is(err, ErrProductNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "booking failed, please try again"})
		}
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetLiveEventsHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	events, err := h.svc.GetLiveEvents(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func (h *Handler) GetEventProductsHandler(c *gin.Context) {
	eventId := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	products, err := h.svc.GetEventProducts(ctx, eventId)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidEventID):
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}