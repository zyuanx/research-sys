package models

type Role struct {
	BaseModel
	Title       string        `gorm:"size:32,not null" json:"title"`
	Desc        string        `gorm:"size:255" json:"desc"`
	Users       []*User       `gorm:"many2many:user_role" json:"-"`
	Permissions []*Permission `gorm:"many2many:role_permission" json:"-"`
}
