package model

type User struct {
	BaseModel
	Username   string  `gorm:"size:20;not null;uniqueIndex" json:"username"`
	Nickname   string  `gorm:"size:20" json:"nickname"`
	Name       string  `gorm:"size:20" json:"name"`
	Gender     string  `gorm:"size:8" json:"gender"`
	College    string  `gorm:"size:32" json:"college"`
	Rank       string  `gorm:"size:16" json:"rank"`
	Profession string  `gorm:"size:32" json:"profession"`
	Clasz      string  `gorm:"size:32" json:"clasz"`
	Typ        string  `gorm:"size:16" json:"typ"`
	Password   string  `gorm:"size:255;not null" json:"-"`
	Telephone  string  `gorm:"size:11;" json:"telephone"`
	Email      string  `gorm:"size:255" json:"email"`
	Roles      []*Role `gorm:"many2many:user_role" json:"-"`
}
