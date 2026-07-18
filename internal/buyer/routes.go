package buyer

import (
	"Orbit/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyerRoutes(r *gin.RouterGroup) {
	r.Use(middleware.UserMiddleware(), middleware.BuyerMiddleware())

	svc := NewService(false, 0) // requirePayment=false until payment gateway is ready
	h := NewHandler(svc)

	r.POST("/event/:eventId/purchase/:productId", h.BuyerEventHandler)

	r.GET("/events", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented yet"})
	})
	r.GET("/event/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented yet"})
	})
}