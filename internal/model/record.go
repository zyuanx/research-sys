package model

type Record struct {
	BaseModel
	Title      string `gorm:"size:32,not null" json:"title"`
	ResearchID string `gorm:"size:128;index" json:"researchID"`
	RecordID   string `gorm:"size:128" json:"recordID"`
	IP         string `gorm:"size:64" json:"ip"`
	Publisher  int    `gorm:"size:64;index" json:"publisher"`
	UserID     int    `json:"-"`
	User       User   `json:"user"`
}

type RecordMgo struct {
	FieldsValue map[string]interface{} `json:"fieldsValue" bson:"fieldsValue"`
}
