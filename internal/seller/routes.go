package seller

import (
	"Orbit/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SellerRoutes(r *gin.RouterGroup) {
	r.Use(
		middleware.UserMiddleware(),
		middleware.SellerMiddleware(),
	)
}