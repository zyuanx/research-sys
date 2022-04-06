package request

import "time"

type ResearchUpdatePayload struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Config      string    `json:"config" binding:"required"`
	StartAt     time.Time `json:"startAt" binding:"required"`
	EndAt       time.Time `json:"endAt" binding:"required"`
	Access      string    `json:"access" binding:"required"`
	Once        *int      `json:"once" binding:"required"`
}
