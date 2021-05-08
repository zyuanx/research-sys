package models

type Record struct {
	BaseModel
	Title      string `gorm:"size:32,not null" json:"title"`
	ResearchID string `gorm:"size:128;index" json:"researchID"`
	RecordId   string `gorm:"size:128,index" json:"recordId"`
	UserID     int
	User       User
}

type RecordMgo struct {
	Detail map[string]interface{} `json:"detail" bson:"detail"`
}
