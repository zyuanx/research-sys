package initialize

import (
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateAdmin() {
	userService := service.NewUserService()
	user := model.User{
		Username: "admin",
	}
	if err := userService.FindUserByUsername(&user); err != nil {
		log.Println("err", err)
	}
	if user.ID != 0 {
		log.Println("admin already exists")
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.Roles = []*model.Role{
		{
			Title: "common",
			Desc:  "common",
		},
		{
			Title: "admin",
			Desc:  "superuser",
		},
	}
	if err := userService.Create(&user); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("admin create success")
}
