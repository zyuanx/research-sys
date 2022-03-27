package model

type User struct {
	BaseModel
	Username   string  `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Nickname   string  `gorm:"size:64" json:"nickname"`
	Name       string  `gorm:"size:64" json:"name"`
	Gender     string  `gorm:"size:8" json:"gender"`
	College    string  `gorm:"size:64" json:"college"`
	Rank       string  `gorm:"size:16" json:"rank"`
	Profession string  `gorm:"size:64" json:"profession"`
	Classname  string  `gorm:"size:64" json:"classname"`
	Category   string  `gorm:"size:16;comment:类别" json:"category"`
	Password   string  `gorm:"size:255;not null" json:"-"`
	Telephone  string  `gorm:"size:11" json:"telephone"`
	Email      string  `gorm:"size:255" json:"email"`
	Roles      []*Role `gorm:"many2many:user_role" json:"roles"`
}
