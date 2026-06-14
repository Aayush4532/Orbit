package product

import (
	"Orbit/internal/utils"

	"github.com/gin-gonic/gin"
)

func RegisterProductsHandler(c *gin.Context) {
	claim, ok := c.Get("UserFields")
	if !ok {
		c.JSON(400, gin.H{"error": "Failed to retrieve user fields"})
		return
	}

	sellerClaim, ok := claim.(*utils.Claims)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user claims"})
		return
	}

	eventIdStr := c.Param("eventId")

	var req RegisterProductsRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := RegisterProductsService(c.Request.Context(), sellerClaim.ID, eventIdStr, req.Products)
	if err != nil {
		if err.Error() == "event not found" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "unauthorized access to this event" {
			c.JSON(403, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Products registered successfully"})
}

func GetAllEventProductsHandler(c *gin.Context) {
	eventIdStr := c.Param("eventId")

	products, err := GetAllEventProductsService(c.Request.Context(), eventIdStr)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"products": products})
}

func GetAnEventProductHandler(c *gin.Context) {
	productIdStr := c.Param("id")

	product, err := GetAnEventProductService(c.Request.Context(), productIdStr)
	if err != nil {
		if err.Error() == "product not found" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"product": product})
}

func UpdateAnEventProductHandler(c *gin.Context) {
	claim, ok := c.Get("UserFields")
	if !ok {
		c.JSON(400, gin.H{"error": "Failed to retrieve user fields"})
		return
	}

	sellerClaim, ok := claim.(*utils.Claims)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user claims"})
		return
	}

	eventIdStr := c.Param("eventId")
	productIdStr := c.Param("id")

	var req UpdateProductRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := UpdateAnEventProductService(c.Request.Context(), sellerClaim.ID, eventIdStr, productIdStr, req)
	if err != nil {
		if err.Error() == "event not found" || err.Error() == "product not found" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "unauthorized access to this event" {
			c.JSON(403, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product updated successfully"})
}

func DeleteAnEventProductHandler(c *gin.Context) {
	claim, ok := c.Get("UserFields")
	if !ok {
		c.JSON(400, gin.H{"error": "Failed to retrieve user fields"})
		return
	}

	sellerClaim, ok := claim.(*utils.Claims)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid user claims"})
		return
	}

	eventIdStr := c.Param("eventId")
	productIdStr := c.Param("id")

	err := DeleteAnEventProductService(c.Request.Context(), sellerClaim.ID, eventIdStr, productIdStr)
	if err != nil {
		if err.Error() == "event not found" || err.Error() == "product not found" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "unauthorized access to this event" {
			c.JSON(403, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}