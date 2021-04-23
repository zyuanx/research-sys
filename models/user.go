package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;unique;uniqueIndex"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}


//func (u *User) Login() error {
//	password := u.Password
//	result := global.Mysql.Where("username = ?", u.Username).First(&u)
//	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
//	if err != nil {
//		return err
//	}
//	return result.Error
//}
//func (u *User) Add() error {
//	result := global.Mysql.Create(u)
//	return result.Error
//}
