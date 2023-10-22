package model

type Record struct {
	BaseModel
	ResearchID int      `gorm:"comment:问卷ID" json:"-"`
	Research   Research `gorm:"foreignKey:ResearchID" json:"research"`
	IPAddress  string   `gorm:"size:64;comment:填写IP" json:"IPAddress"`
	Status     int      `gorm:"default:0;comment:审核状态" json:"status"`
	Values     string   `gorm:"type:text;comment:填写者ID" json:"values"`
	UserID     int      `gorm:"comment:填写者ID" json:"userID"`
	User       User     `gorm:"foreignKey:UserID" json:"user"`
}
