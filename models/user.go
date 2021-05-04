package models

type User struct {
	BaseModel
	Username  string  `gorm:"size:20;not null;uniqueIndex" json:"username"`
	Nickname  string  `gorm:"size:20;" json:"nickname"`
	Password  string  `gorm:"size:255;not null" json:"-"`
	Avatar    string  `gorm:"size:255;" json:"avatar"`
	Telephone string  `gorm:"size:11;" json:"telephone"`
	Email     string  `gorm:"size:255" json:"email"`
	Roles     []*Role `gorm:"many2many:user_role" json:"roles"`
}
