package models

import (
	"gin-research-sys/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//func init()  {
//	if err := common.DB.AutoMigrate(&User{}); err != nil {
//		log.Println(err)
//	}
//}
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;unique;uniqueIndex"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

type Users []User

func (u *User) Login() error {
	password := u.Password
	result := common.DB.Where("username = ?", u.Username).First(&u)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return result.Error
}
func (u *User) Add() error {
	result := common.DB.Create(u)
	return result.Error
}
