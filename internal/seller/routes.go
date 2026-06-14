package seller

import (
	"Orbit/internal/middleware"
	"Orbit/internal/seller/verify"
	"Orbit/internal/seller/isVerified"
	"github.com/gin-gonic/gin"
	"Orbit/internal/seller/event"
)



func SellerRoutes(r *gin.RouterGroup) {
	r.Use(
		middleware.UserMiddleware(),
		middleware.SellerMiddleware(),
	)

	r.GET("/isVerified", IsVerified.IsVerifiedSeller);
	r.POST("/verify", verify.VerifySeller);

	EventGroups := r.Group("/events");
	event.EventRoutes(EventGroups);
}