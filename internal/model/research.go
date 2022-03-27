package model

import "time"

type Research struct {
	BaseModel
	Title       string    `gorm:"size:64,not null;comment:标题" json:"title"`
	Desc        string    `gorm:"size:255;comment:描述" json:"desc"`
	Rules       string    `gorm:"type:text;comment:数据校验" json:"rules"`
	Fields      string    `gorm:"type:text;comment:问题字段" json:"fieldsValue"`
	StartAt     time.Time `gorm:"type:datetime;index;comment:开始时间" json:"startAt"`
	EndAt       time.Time `gorm:"type:datetime;index;comment:结束时间" json:"endAt"`
	Access      string    `gorm:"size:64;index;comment:访问者" json:"access"`
	Status      int       `gorm:"default:0" json:"status"`
	Once        int       `gorm:"default:0;comment:是否仅可填写一次" json:"once"`
	PublisherID int       `gorm:"comment:发布者ID" json:"-"`
	Publisher   User      `gorm:"foreignKey:PublisherID" json:"publisher"`
}
