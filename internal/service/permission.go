package service

import (
	"github.com/zyuanx/research-sys/internal/model"

	"github.com/zyuanx/research-sys/internal/pkg/pagination"
)

func (s *Service) PermissionList(page int, size int, permissions *[]model.Permission, total *int64) error {
	if err := s.db.Model(&model.Permission{}).Count(total).
		Scopes(pagination.Paginate(page, size)).
		Order("`permissions`.`group`").Order("`permissions`.`index`").
		Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) PermissionRetrieve(permission *model.Permission, id int) error {
	if err := s.db.Model(&model.Permission{}).First(&permission, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) PermissionCreate(permission *model.Permission) error {
	if err := s.db.Model(&model.Permission{}).Create(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) PermissionUpdate(permission *model.Permission) error {
	if err := s.db.Save(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) PermissionDestroy(id int) error {
	if err := s.db.Delete(&model.Permission{}, id).Error; err != nil {
		return err
	}
	return nil
}
