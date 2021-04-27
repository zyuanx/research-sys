package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IUserService interface {
	UserRegister(user *models.User) error
	UserLogin(user *models.User) error
	UserInfo(user *models.User, id uint64) error
}
type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (u UserService) UserRegister(user *models.User) error {
	if err := global.Mysql.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) UserLogin(user *models.User) error {

	result := global.Mysql.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserService) UserInfo(user *models.User, id uint64) error {
	result := global.Mysql.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
