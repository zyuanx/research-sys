package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IRoleService interface {
	Create(role *models.Role) error
}

func (r RoleService) Create(role *models.Role) error {
	if err := global.Mysql.Create(role).Error; err != nil {
		return err
	}
	return nil
}

type RoleService struct {
}

func NewRoleService() RoleService {
	return RoleService{}
}
