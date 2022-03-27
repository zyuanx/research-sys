package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/util"
)

type IUserService interface {
	FindByUsername(user *model.User, username string) error

	List(users *[]model.User, page int, size int, total *int64, query map[string]interface{}) error
	Retrieve(user *model.User, id int) error
	Create(user *model.User) error
	Update(user *model.User, payload map[string]interface{}) error
	UpdateByUser(user *model.User, patch map[string]interface{}) (err error)
	Destroy(id int) error
	UpdateRole(user *model.User, rolesID []int) error
}
type UserService struct{}

func NewUserService() IUserService {
	return UserService{}
}

func (u UserService) FindByUsername(user *model.User, username string) error {
	if err := conf.Mysql.Model(&model.User{}).
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) List(users *[]model.User, page int, size int, total *int64, query map[string]interface{}) error {
	if err := conf.Mysql.Model(&model.User{}).
		Where(query).Count(total).
		Scopes(util.Paginate(page, size)).
		Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Retrieve(user *model.User, id int) error {
	if err := conf.Mysql.Model(&model.User{}).
		Preload("Roles").
		First(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Create(user *model.User) error {
	if err := conf.Mysql.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Update(user *model.User, payload map[string]interface{}) error {
	if err := conf.Mysql.Model(&user).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) UpdateByUser(user *model.User, patch map[string]interface{}) error {
	if err := conf.Mysql.Model(&user).Updates(patch).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Destroy(id int) error {
	if err := conf.Mysql.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) UpdateRole(user *model.User, rolesID []int) error {
	var roles []model.Role
	if err := conf.Mysql.Model(&model.Role{}).Find(&roles, "id IN ?", rolesID).Error; err != nil {
		return err
	}
	if err := conf.Mysql.Model(&user).Association("Roles").Replace(roles); err != nil {
		return err
	}
	return nil
}
