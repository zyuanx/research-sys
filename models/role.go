package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title       string        `gorm:"size:32,not null"`
	Desc        string        `gorm:"size:255"`
	Users       []*User       `gorm:"many2many:user_role"`
	Permissions []*Permission `gorm:"many2many:role_permission"`
}
