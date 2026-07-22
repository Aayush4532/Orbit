package buyer

import (
	"Orbit/internal/middleware"

	"github.com/gin-gonic/gin"
)

func BuyerRoutes(r *gin.RouterGroup) {
	r.Use(middleware.UserMiddleware(), middleware.BuyerMiddleware())

	svc := NewService(false, 0)
	h := NewHandler(svc)

	r.GET("/events", h.GetLiveEventsHandler)

	r.GET("/event/:id", h.GetEventProductsHandler)

	r.POST("/event/:eventId/purchase/:productId", h.BuyerEventHandler)
}