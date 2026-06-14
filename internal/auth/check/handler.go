package check

import (
	"Orbit/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// the function checks if user is signedin or not
func CheckHandler(c *gin.Context) {
	value, exists := c.Get("UserFields")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}

	claims := value.(*utils.Claims)

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"message":       "user is signedin",
		"user": gin.H{
			"id":              claims.ID,
			"emailId":         claims.EmailId,
			"role":            claims.Role,
			"isEmailVerified": claims.IsEmailVerified,
			"isActive":        claims.IsActive,
		},
	})
}
