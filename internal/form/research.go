package form

import "time"

type ResearchListForm struct {
	Pagination
	Access string `json:"access" binding:"required"`
}
type ResearchCreateForm struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Config      string    `json:"config" binding:"required"`
	Items       string    `json:"items" binding:"required"`
	Values      string    `json:"values" binding:"required"`
	StartAt     time.Time `json:"startAt" binding:"required"`
	EndAt       time.Time `json:"endAt" binding:"required"`
	Access      string    `json:"access" binding:"required"`
	Once        int       `json:"once"`
	Open        int       `json:"open"`
	PublisherID int       `json:"publisherID"`
}
