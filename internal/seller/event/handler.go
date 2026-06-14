package event

import (
	"Orbit/internal/models"
	"Orbit/internal/repositories"
	"Orbit/internal/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateAnEventHandler(c *gin.Context) {
	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	claim, ok := c.Get("UserFields")
	if !ok {
		c.JSON(400, gin.H{"error": "Failed to retrieve user fields"})
		return
	}
	sellerClaim := claim.(*utils.Claims)
	ObjectifiedId, err := utils.GetObjectFiedIdFromString(sellerClaim.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	newEventId := bson.NewObjectID()

	event := &models.Event{
		ID:          newEventId,
		SellerID:    ObjectifiedId,
		EventName:   req.Title,
		Description: req.Description,
		ScheduledAt: req.ScheduledAt,
		ImageBanner: req.ImageBanner,
		IsLive:      false,
	}

	err = repositories.CreateAnEvent(event)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(201, gin.H{"message": "Event created successfully", "eventId": newEventId.Hex()})
}

func GetAllEventsHandler(c *gin.Context) {
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

	events, err := repositories.GetAllEvents(sellerClaim.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(200, gin.H{"events": events})
}

func GetAnEventHandler(c *gin.Context) {
	eventId := c.Param("id")
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

	event, err := repositories.GetAnEvent(eventId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.SellerID.Hex() != sellerClaim.ID {
		c.JSON(403, gin.H{"error": "Unauthorized access to event"})
		return
	}

	c.JSON(200, gin.H{"event": event})
}

func UpdateAnEventHandler(c *gin.Context) {
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

	eventId := c.Param("id")
	
	event, err := repositories.GetAnEvent(eventId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.SellerID.Hex() != sellerClaim.ID {
		c.JSON(403, gin.H{"error": "Unauthorized access to this event"})
		return
	}

	var req UpdateEventRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updateData := bson.M{}
	if req.Title != nil {
		updateData["eventName"] = *req.Title
	}
	if req.Description != nil {
		updateData["description"] = *req.Description
	}
	if req.ScheduledAt != nil {
		updateData["scheduledAt"] = *req.ScheduledAt
	}
	if req.ImageBanner != nil {
		updateData["imageBanner"] = *req.ImageBanner
	}

	if len(updateData) == 0 {
		c.JSON(400, gin.H{"error": "No fields provided for update"})
		return
	}

	if err := repositories.UpdateAnEvent(eventId, updateData); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Event updated successfully"})
}

func DeleteAnEventHandler(c *gin.Context) {
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

	eventId := c.Param("id")
	event, err := repositories.GetAnEvent(eventId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found or does not exist"})
		return
	}

	if event == nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.SellerID.Hex() != sellerClaim.ID {
		c.JSON(403, gin.H{"error": "Unauthorized access to this event"})
		return
	}

	if err := repositories.DeleteAnEvent(eventId); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(200, gin.H{"message": "Event deleted successfully"})
}