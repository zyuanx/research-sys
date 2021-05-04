package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
	"gin-research-sys/pkg/utils"
)

type IRoleService interface {
	List(page int, size int, roles *[]models.Role, total *int64)
	Retrieve(role *models.Role, id uint) error
	Create(role *models.Role) error
	Update(role *models.Role, id uint, data interface{}) error
	Destroy(role *models.Role, id uint) error
}

type RoleService struct {
}

func NewRoleService() RoleService {
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

func (r RoleService) Retrieve(role interface{}, id int) error {
	if err := global.Mysql.Model(&models.Role{}).First(&role, id).Error; err != nil {
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

func (r RoleService) Update(role *models.Role, data interface{}) error {
	d, _ := utils.ToMap(data, "json")
	if err := global.Mysql.Model(&role).Updates(&d).Error; err != nil {
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
