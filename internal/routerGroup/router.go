package routergroup

import (
	"Orbit/internal/auth"
	"Orbit/internal/seller"

	"github.com/gin-gonic/gin"
)

func ApiRoutes (r *gin.RouterGroup) {
	// declare all route groups.
	AuthRouterGroup := r.Group("/auth");
	SellerRouterGroup := r.Group("/seller");
	// implement grouping.
	auth.AuthRoutes(AuthRouterGroup);
	seller.SellerRoutes(SellerRouterGroup);
}