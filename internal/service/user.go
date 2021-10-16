package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/util"
)

type IUserService interface {
	FindUserByUsername(user *model.User) error

	List(page int, size int, users *[]model.User, total *int64) error
	Retrieve(user *model.User, id int) error
	Create(user *model.User) error
	Update(user *model.User) error
	Destroy(id int) error

	ListRole2(user *model.User, roles *[]int) error
	UpdateRole(user *model.User, ids []int) error
}
type UserService struct{}

func NewUserService() IUserService {
	return UserService{}
}

func (u UserService) FindUserByUsername(user *model.User) error {
	if err := conf.Mysql.
		Where("username = ?", user.Username).
		First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) List(page int, size int, users *[]model.User, total *int64) error {
	if err := conf.Mysql.Model(&model.User{}).Count(total).
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

func (u UserService) Update(user *model.User) error {
	if err := conf.Mysql.Save(&user).Error; err != nil {
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

func (u UserService) ListRole2(user *model.User, roles *[]int) error {
	var t []int
	for _, value := range user.Roles {
		t = append(t, int(value.ID))
	}
	*roles = t
	return nil
}

func (u UserService) UpdateRole(user *model.User, ids []int) error {
	var roles []model.Role
	if err := conf.Mysql.Model(&model.Role{}).Find(&roles, "id IN ?", ids).Error; err != nil {
		return err
	}
	if err := conf.Mysql.Model(&user).Association("Roles").Replace(roles); err != nil {
		return err
	}
	return nil
}
