package initialize

import (
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/service"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateAdmin() {
	userService := service.NewUserService()
	user := model.User{}
	if err := userService.FindByUsername(&user, "admin"); err != nil {
		log.Println("err", err)
		return
	}
	if user.ID != 0 {
		log.Println("admin already exists")
		return
	}
	user.Username = "admin"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.Roles = []*model.Role{
		{
			Title:       "admin",
			Description: "admin",
		},
		{
			Title:       "super",
			Description: "super",
		},
	}
	if err := userService.Create(&user); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("admin create success")
}
