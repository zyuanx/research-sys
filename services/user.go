package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IUserService interface {
	UserLogin(user *models.User) error
	UserInfo(user *models.User) error
	UserPasswordReset(user *models.User) error

	List(page int, size int, users *[]models.User, total *int64) error
	Retrieve(user *models.User, id int) error
	Create(user *models.User) error
	Update(user *models.User) error
	Destroy(id int) error
}
type UserService struct{}

func NewUserService() IUserService {
	return UserService{}
}

func (u UserService) UserLogin(user *models.User) error {

	result := global.Mysql.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserService) UserInfo(user *models.User) error {
	if err := global.Mysql.Model(&models.User{}).
		Preload("Roles").First(&user).Error; err != nil {
		return err
	}
	return nil
}
func (u UserService) UserPasswordReset(user *models.User) error {
	panic("implement me")
}
func (u UserService) List(page int, size int, users *[]models.User, total *int64) error {
	if err := global.Mysql.Model(&models.User{}).Count(total).
		Scopes(global.Paginate(page, size)).
		Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Retrieve(user *models.User, id int) error {
	if err := global.Mysql.Model(&models.User{}).First(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Create(user *models.User) error {
	if err := global.Mysql.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Update(user *models.User) error {
	if err := global.Mysql.Omit("username").Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Destroy(id int) error {
	if err := global.Mysql.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
