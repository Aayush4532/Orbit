package event

import "time"

type RequestBody struct {
	Title       string    `json:"title" binding:"required,min=3,max=100"`
	Description string    `json:"description" binding:"required,min=10,max=500"`
	ScheduledAt time.Time `json:"scheduledAt" binding:"required"`
	ImageBanner string    `json:"imageBanner" binding:"required"`
}

type UpdateEventRequestBody struct {
	Title       *string    `json:"title" binding:"omitempty,min=3,max=100"`
	Description *string    `json:"description" binding:"omitempty,min=10,max=500"`
	ScheduledAt *time.Time `json:"scheduledAt" binding:"omitempty"`
	ImageBanner *string    `json:"imageBanner" binding:"omitempty"`
}