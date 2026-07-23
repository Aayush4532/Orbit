package event

import "time"

type RequestBody struct {
	Title       string    `form:"title"       binding:"required"`
	Description string    `form:"description" binding:"required"`
	ScheduledAt time.Time `form:"scheduledAt" binding:"required"`
}

type UpdateEventRequestBody struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	ScheduledAt *time.Time `json:"scheduledAt"`
	ImageBanner *string    `json:"imageBanner"`
}