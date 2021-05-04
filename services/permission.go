package services

import (
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IPermissionService interface {
	List(page int, size int, permissions *[]models.Permission, total *int64)
	Retrieve(permission *models.Permission, id int) error
	Create(permission *models.Permission) error
	Update(permission *models.Permission, id int, data interface{}) error
	Destroy(permission *models.Permission, id int) error
}
type PermissionService struct {
}

func NewPermissionService() PermissionService {
	return PermissionService{}
}

func (p PermissionService) List(page int, size int, permissions *[]models.Permission, total *int64) error {
	if err := global.Mysql.Model(&models.Permission{}).Count(total).
		Scopes(global.Paginate(page, size)).
		Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func (p PermissionService) Retrieve(permission *models.Permission, id int) error {
	if err := global.Mysql.Model(&models.Permission{}).First(&permission, id).Error; err != nil {
		return err
	}
	return nil
}

func (p PermissionService) Create(permission *models.Permission) error {
	panic("implement me")
}

func (p PermissionService) Update(permission *models.Permission, id int, data interface{}) error {
	panic("implement me")
}

func (p PermissionService) Destroy(permission *models.Permission, id int) error {
	panic("implement me")
}
