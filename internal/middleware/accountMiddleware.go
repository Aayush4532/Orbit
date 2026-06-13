package middleware

import (
	"Orbit/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(401, gin.H{"error": "Token is not present"})
			c.Abort()
			return
		}

		userData, err := utils.ValidateJwtToken(token);
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H {
				"error" : "session expired",
			})
			c.Abort();
			return
		}
		
		c.Set("UserFields", userData);
		c.Next();
	}
}

func SellerMiddleware(claim *utils.Claims) bool { // just a func to check if the user is seller but not actual middleware implement. 
	if claim.Role != "seller" {
		return false;
	} 
	return true;
}