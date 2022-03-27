package model

type Record struct {
	BaseModel
	ResearchID  int      `gorm:"comment:问卷ID" json:"-"`
	Research    Research `gorm:"foreignKey:ResearchID" json:"research"`
	IP          string   `gorm:"size:64;comment:填写IP" json:"ip"`
	Status      int      `gorm:"size:8;index;comment:审核状态" json:"status"`
	FieldsValue string   `gorm:"type:text" bson:"fieldsValue"`
	WriterID    int      `gorm:"comment:填写者ID" json:"-"`
	Writer      User     `json:"writer"`
}
