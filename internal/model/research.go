package model

import "time"

type Research struct {
	BaseModel
	Title       string    `gorm:"size:64,not null;comment:标题" json:"title"`
	Description string    `gorm:"size:255;comment:描述" json:"description"`
	Config      string    `gorm:"type:text;comment:配置" json:"config"`
	Items       string    `gorm:"type:text;comment:字段" json:"items"`
	Values      string    `gorm:"type:text;comment:值" json:"values"`
	StartAt     time.Time `gorm:"type:datetime;comment:开始时间" json:"startAt"`
	EndAt       time.Time `gorm:"type:datetime;comment:结束时间" json:"endAt"`
	Access      string    `gorm:"size:255;index;comment:访问者" json:"access"`
	Once        int       `gorm:"default:0;comment:重复填写" json:"once"`
	Open        int       `gorm:"default:0;comment:是否公开" json:"open"`
	PublisherID int       `gorm:"comment:发布者ID" json:"publisherID"`
	Publisher   User      `gorm:"foreignKey:PublisherID" json:"publisher"`
}
