package sale

import (
	"Orbit/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LiveSaleHandler (c *gin.Context) {
	eventId := c.Param("eventId");
	UserClaim, ok := c.Get("UserFields");
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"error" : "not authorized to perform this operation",
		})
		return
	}

	claim := UserClaim.(*utils.Claims);
	sellerId := claim.ID;

	err := LiveSaleService(sellerId, eventId);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "sale live successfully",
	})
}

// func PauseSaleHandler (c *gin.Context) {
// 	eventId := c.Param("eventId");
// 	UserClaim, ok := c.Get("UserFields");
// 	if !ok {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error" : "not authorized to perform this operation",
// 		})
// 	}

// 	claim := UserClaim.(*utils.Claims);
// 	if err != nil {
// 		c.JSON(http.status)
// 	}
// }

// func StopSaleHandler (c *gin.Context) {
// 	eventId := c.Param("eventId");	
// 	UserClaim, ok := c.Get("UserFields");
// 	if !ok {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error" : "not authorized to perform this operation",
// 		})
// 	}

// 	claim := UserClaim.(*utils.Claims);
// 	if err != nil {
// 		c.JSON(http.status)
// 	}
// }