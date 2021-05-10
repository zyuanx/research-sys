package models

type Research struct {
	BaseModel
	Title      string `gorm:"size:32,not null" json:"title"`
	Desc       string `gorm:"size:255" json:"desc"`
	Status     int    `gorm:"default:0" json:"status"`
	Once       int    `gorm:"default:0" json:"once"`
	ResearchID string `gorm:"size:128;index" json:"researchID"`
	UserID     int    `json:"-"`
	User       User   `json:"-"`
}

type ResearchMgo struct {
	FieldsValue map[string]interface{}   `json:"fieldsValue" bson:"fieldsValue" binding:"required"`
	Detail      []map[string]interface{} `json:"detail" bson:"detail" binding:"required"`
	Rules       map[string]interface{}   `json:"rules" bson:"rules" binding:"required"`
}
