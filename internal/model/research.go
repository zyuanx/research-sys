package model

import (
	"encoding/json"
	"time"
)

type Research struct {
	BaseModel
	Title       string    `gorm:"size:64,not null;comment:标题" json:"title"`
	Description string    `gorm:"size:255;comment:描述" json:"description"`
	Config      string    `gorm:"type:text;comment:配置" json:"config"`
	Pattern     string    `gorm:"type:text;comment:样式" json:"pattern"`
	Items       string    `gorm:"type:text;comment:字段" json:"items"`
	StartAt     time.Time `gorm:"type:datetime;comment:开始时间" json:"startAt"`
	EndAt       time.Time `gorm:"type:datetime;comment:结束时间" json:"endAt"`
	Once        int       `gorm:"default:0;comment:重复填写" json:"once"`
	Open        int       `gorm:"default:0;comment:是否公开" json:"open"`
	PublisherID uint      `gorm:"comment:发布者ID" json:"publisherID"`
	Publisher   User      `gorm:"foreignKey:PublisherID" json:"publisher"`
}

type ResearchListReq struct {
	Pagination
}

type ResearchCreateReq struct {
	Title       string                   `json:"title" binding:"required"`
	Description string                   `json:"description"`
	Config      map[string]interface{}   `json:"config" binding:"required"`
	Pattern     map[string]interface{}   `json:"pattern" binding:"required"`
	Items       []map[string]interface{} `json:"items" binding:"required"`
	StartAt     time.Time                `json:"startAt" binding:"required"`
	EndAt       time.Time                `json:"endAt" binding:"required"`
	Once        *int                     `json:"once" binding:"required"`
	Open        *int                     `json:"open" binding:"required"`
}

type ResearchUpdateReq struct {
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Config      map[string]interface{}   `json:"config"`
	Pattern     map[string]interface{}   `json:"pattern"`
	Items       []map[string]interface{} `json:"items"`
	StartAt     time.Time                `json:"startAt"`
	EndAt       time.Time                `json:"endAt"`
	Once        *int                     `json:"once"`
	Open        *int                     `json:"open"`
}

type ResearchRes struct {
	ID          uint                     `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Config      map[string]interface{}   `json:"config" binding:"required"`
	Pattern     map[string]interface{}   `json:"pattern" binding:"required"`
	Items       []map[string]interface{} `json:"items" binding:"required"`
	StartAt     time.Time                `json:"startAt" binding:"required"`
	EndAt       time.Time                `json:"endAt" binding:"required"`
	Once        int                      `json:"once" binding:"required"`
	Open        int                      `json:"open" binding:"required"`
	PublisherID uint                     `json:"publisherID"`
	CreatedAt   time.Time                `json:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt"`
}

func (r Research) ToRes() ResearchRes {
	res := ResearchRes{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		StartAt:     r.StartAt,
		EndAt:       r.EndAt,
		Once:        r.Once,
		Open:        r.Open,
		PublisherID: r.PublisherID,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
	config := make(map[string]interface{})
	if err := json.Unmarshal([]byte(r.Config), &config); err != nil {
		return res
	}
	res.Config = config
	pattern := make(map[string]interface{})
	if err := json.Unmarshal([]byte(r.Pattern), &pattern); err != nil {
		return res
	}
	res.Pattern = pattern
	items := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(r.Items), &items); err != nil {
		return res
	}
	res.Items = items
	return res

}
