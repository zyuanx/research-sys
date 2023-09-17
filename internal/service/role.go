package service

import (
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/pagination"
)

func (s *Service) RoleList(roles *[]model.Role, page int, size int, total *int64, query map[string]interface{}) error {
	if err := s.db.Model(&model.Role{}).
		Where(query).Count(total).
		Scopes(pagination.Paginate(page, size)).
		Find(&roles).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RoleRetrieve(role *model.Role, id int) error {
	if err := s.db.Model(&model.Role{}).
		Preload("Permissions").
		First(&role, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RoleCreate(role *model.Role) error {
	if err := s.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RoleUpdate(role *model.Role, payload map[string]interface{}) error {
	if err := s.db.Model(&role).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RoleDelete(id int) error {
	if err := s.db.Delete(&model.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

// func (s *Service) RoleUpdatePermission(role *model.Role, ids []int) error {
// 	var permissions []model.Permission
// 	if err := s.db.Model(&model.Permission{}).
// 		Find(&permissions, "id IN ?", ids).Error; err != nil {
// 		return err
// 	}
// 	if err := s.db.Model(&role).
// 		Association("Permissions").
// 		Replace(permissions); err != nil {
// 		return err
// 	}
// 	return nil
// }
