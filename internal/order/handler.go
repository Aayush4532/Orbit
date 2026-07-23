package order

import (
	"Orbit/internal/repositories"
	"Orbit/internal/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetMyOrdersHandler(c *gin.Context) {
	raw, ok := c.Get("UserFields")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "not permitted"})
		return
	}
	claim, ok := raw.(*utils.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid session"})
		return
	}

	userId, err := bson.ObjectIDFromHex(claim.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	orders, err := repositories.GetOrdersByUser(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetSellerEventOrdersHandler(c *gin.Context) {
	eventIdStr := c.Param("eventId")

	eventId, err := bson.ObjectIDFromHex(eventIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	orders, err := repositories.GetOrdersByEvent(ctx, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetEventAnalyticsHandler(c *gin.Context) {
	eventIdStr := c.Param("eventId")

	eventId, err := bson.ObjectIDFromHex(eventIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	analytics, err := repositories.GetEventAnalytics(ctx, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve analytics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"analytics": analytics})
}
