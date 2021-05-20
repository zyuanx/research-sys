package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IRoleService interface {
	List(page int, size int, roles *[]models.Role, total *int64) error
	Retrieve(role *models.Role, id int) error
	Create(role *models.Role) error
	Update(role *models.Role) error
	Destroy(id int) error

	UpdatePermission(role *models.Role, ids []int) error
}

type RoleService struct{}

func NewRoleService() IRoleService {
	return RoleService{}
}

func (r RoleService) List(page int, size int, roles *[]models.Role, total *int64) error {
	if err := global.Mysql.Model(&models.Role{}).Count(total).
		Scopes(global.Paginate(page, size)).
		Find(&roles).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Retrieve(role *models.Role, id int) error {
	if err := global.Mysql.Model(&models.Role{}).
		Preload("Permissions").
		First(&role, id).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Create(role *models.Role) error {
	if err := global.Mysql.Model(&models.Role{}).Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Update(role *models.Role) error {
	if err := global.Mysql.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r RoleService) Destroy(id int) error {
	if err := global.Mysql.Delete(&models.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}


func (r RoleService) UpdatePermission(role *models.Role, ids []int) error {
	var permissions []models.Permission
	if err := global.Mysql.Model(&models.Permission{}).
		Find(&permissions, "id IN ?", ids).Error; err != nil {
		return err
	}
	if err := global.Mysql.Model(&role).
		Association("Permissions").
		Replace(permissions); err != nil {
		return err
	}
	return nil
}