package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IRoleService interface {
	List(roles []*models.Role, size int, page uint) error
	Create(role *models.Role) error
}

type RoleService struct {
}

func NewRoleService() RoleService {
	return RoleService{}
}

func (r RoleService) List(roles *[]models.Role, size int, page uint) error {
	if err := global.Mysql.Limit(size).Find(&roles).Error; err != nil {
		return err
	}
	return nil
}
func (r RoleService) Create(role *models.Role) error {
	if err := global.Mysql.Create(role).Error; err != nil {
		return err
	}
	return nil
}
