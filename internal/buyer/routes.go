package buyer

import (
	"Orbit/internal/middleware"

	"github.com/gin-gonic/gin"
)

func BuyerRoutes(r *gin.RouterGroup) {
	r.Use(middleware.UserMiddleware(), middleware.BuyerMiddleware())
	service := NewService(false, 0)
	buyerHandler := NewHandler(service)
	r.GET("/events")
	r.GET("/event/:Id")
	r.POST("/event/:eventId/purchase/:productId", buyerHandler.BuyerEventHandler)
}
