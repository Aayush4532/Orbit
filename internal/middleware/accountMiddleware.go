package middleware

import (
	"Orbit/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(401, gin.H{"error": "Token is not present"})
			c.Abort()
			return
		}

		userData, err := utils.ValidateJwtToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "session expired",
			})
			c.Abort()
			return
		}

		c.Set("UserFields", userData)
		c.Next()
	}
}

func SellerMiddleware() gin.HandlerFunc { // just a func to check if the user is seller but not actual middleware implement.
	return func(c *gin.Context) {

		claims, ok := c.Get("UserFields")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		claim, ok := claims.(*utils.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Invalid user data",
			})
			return
		}

		if claim.Role != "seller" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Seller access required",
			})
			return
		}

		c.Next()
	}
}
