package service

import (
	"github.com/zyuanx/research-sys/internal/model"

	"github.com/zyuanx/research-sys/internal/pkg/pagination"
)

func (s *Service) ListPermission(page int, size int, permissions *[]model.Permission, total *int64) error {
	if err := s.db.Model(&model.Permission{}).Count(total).
		Scopes(pagination.Paginate(page, size)).
		Order("`permissions`.`group`").Order("`permissions`.`index`").
		Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RetrievePermission(permission *model.Permission, id int) error {
	if err := s.db.Model(&model.Permission{}).First(&permission, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) CreatePermission(permission *model.Permission) error {
	if err := s.db.Model(&model.Permission{}).Create(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdatePermission(permission *model.Permission) error {
	if err := s.db.Save(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DestroyPermission(id int) error {
	if err := s.db.Delete(&model.Permission{}, id).Error; err != nil {
		return err
	}
	return nil
}
