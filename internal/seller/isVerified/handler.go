package IsVerified

import (
	"Orbit/internal/repositories"
	"Orbit/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsVerifiedSeller(c *gin.Context) {
	value, exists := c.Get("UserFields")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	claim := value.(*utils.Claims)
	emailId := claim.EmailId

	user, err := repositories.GetUserFromEmail(c, emailId)
	if err != nil {
		if err.Error() == "No Email Exists" {
			c.JSON(http.StatusNotFound, gin.H {
				"error" : err.Error(),
			})
			return;
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	info := user.SellerInfo;
	if info == nil {
		c.JSON(http.StatusForbidden, gin.H {
			"error" : "user is not a seller",
		})
		return
	}

	if info.IsApproved == false {
		c.JSON(http.StatusForbidden, gin.H{
			"error" : "seller is not verified",
			"next_allowed_date" : info.NextAllowedAt,
		})
		return;
	}


	c.JSON(http.StatusOK, gin.H {
		"verified" : true,
		"message" : "Seller is Verified",
	})
}
