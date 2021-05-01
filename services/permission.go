package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IPermissionService interface {
	List(permissions []*models.Permission) error
	Create(permission *models.Permission) error
}

func (r PermissionService) List(permissions []*models.Permission, size int, page uint) error {
	if err := global.Mysql.Limit(size).Offset(5).Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func (r PermissionService) Create(role *models.Permission) error {
	if err := global.Mysql.Create(role).Error; err != nil {
		return err
	}
	return nil
}

type PermissionService struct {
}

func NewPermissionService() PermissionService {
	return PermissionService{}
}
