package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/util"
)

type IRoleService interface {
	List(page int, size int, roles *[]model.Role, total *int64) error
	Retrieve(role *model.Role, id int) error
	Create(role *model.Role) error
	Update(role *model.Role) error
	Destroy(id int) error

	UpdatePermission(role *model.Role, ids []int) error
}

type RoleService struct{}

func NewRoleService() IRoleService {
	return RoleService{}
}

func (r RoleService) List(page int, size int, roles *[]model.Role, total *int64) error {
	if err := conf.Mysql.Model(&model.Role{}).Count(total).
		Scopes(util.Paginate(page, size)).
		Find(&roles).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Retrieve(role *model.Role, id int) error {
	if err := conf.Mysql.Model(&model.Role{}).
		Preload("Permissions").
		First(&role, id).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Create(role *model.Role) error {
	if err := conf.Mysql.Model(&model.Role{}).Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Update(role *model.Role) error {
	if err := conf.Mysql.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Destroy(id int) error {
	if err := conf.Mysql.Delete(&model.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}


func (r RoleService) UpdatePermission(role *model.Role, ids []int) error {
	var permissions []model.Permission
	if err := conf.Mysql.Model(&model.Permission{}).
		Find(&permissions, "id IN ?", ids).Error; err != nil {
		return err
	}
	if err := conf.Mysql.Model(&role).
		Association("Permissions").
		Replace(permissions); err != nil {
		return err
	}
	return nil
}