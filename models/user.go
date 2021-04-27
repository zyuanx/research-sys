package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"size:20;not null;uniqueIndex"`
	Password  string  `gorm:"size:255;not null"`
	Telephone string  `gorm:"size:11;"`
	Avatar    string  `gorm:"size:255;"`
	Roles     []*Role `gorm:"many2many:user_role"`
}
