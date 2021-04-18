package user

import (
	"gin-research-sys/common/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	db = database.InitDB()
	err := db.AutoMigrate(User{})
	if err != nil {

		log.Println("migrate error", err)
		panic("migrate error")
	}
}

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

type Users []User

func (u *User) login() error {
	password := u.Password
	result := db.Where("username = ?", u.Username).First(&u)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return result.Error
}
func (u *User) add() error {
	result := db.Create(u)
	return result.Error
}
