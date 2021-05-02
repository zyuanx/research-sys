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

func (r RoleService) List(page int, size int, modelList interface{}, total *int64) error {
	var err error
	if err = global.Mysql.Model(&models.Role{}).Count(total).Scopes(global.Paginate(page, size)).Find(modelList).Error; err != nil {
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
