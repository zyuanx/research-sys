package model

type Record struct {
	BaseModel
	ResearchID int      `gorm:"comment:问卷ID" json:"-"`
	Research   Research `gorm:"foreignKey:ResearchID" json:"research"`
	IPAddress  string   `gorm:"size:64;comment:填写IP" json:"IPAddress"`
	Status     int      `gorm:"default:0;comment:审核状态" json:"status"`
	Values     string   `gorm:"type:text;comment:填写者ID" json:"values"`
	WriterInfo string   `gorm:"type:text;comment:填写者用户信息" json:"-"`
}

type OpenRecord struct {
	BaseModel
	ResearchID int      `gorm:"comment:问卷ID" json:"-"`
	Research   Research `gorm:"foreignKey:ResearchID" json:"-"`
	IPAddress  string   `gorm:"size:64;comment:填写IP" json:"IPAddress"`
	Values     string   `gorm:"type:text;comment:填写者ID" json:"values"`
}
