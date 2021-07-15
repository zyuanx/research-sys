package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/util"
)

type IPermissionService interface {
	List(page int, size int, permissions *[]model.Permission, total *int64) error
	Retrieve(permission *model.Permission, id int) error
	Create(permission *model.Permission) error
	Update(permission *model.Permission) error
	Destroy(id int) error
}
type PermissionService struct{}

func NewPermissionService() IPermissionService {
	return PermissionService{}
}

func (p PermissionService) List(page int, size int, permissions *[]model.Permission, total *int64) error {
	if err := conf.Mysql.Model(&model.Permission{}).Count(total).
		Scopes(util.Paginate(page, size)).
		Order("`permissions`.`group`").Order("`permissions`.`index`").
		Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func (p PermissionService) Retrieve(permission *model.Permission, id int) error {
	if err := conf.Mysql.Model(&model.Permission{}).First(&permission, id).Error; err != nil {
		return err
	}
	return nil
}

func (p PermissionService) Create(permission *model.Permission) error {
	if err := conf.Mysql.Model(&model.Permission{}).Create(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (p PermissionService) Update(permission *model.Permission) error {
	if err := conf.Mysql.Save(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (p PermissionService) Destroy(id int) error {
	if err := conf.Mysql.Delete(&model.Permission{}, id).Error; err != nil {
		return err
	}
	return nil
}
