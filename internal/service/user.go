package service

import (
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/pagination"
)

func (s *Service) UserFindByUsername(user *model.User, username string) error {
	if err := s.db.Model(&model.User{}).
		Where("username = ?", username).
		Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UserList(users *[]model.User, page int, size int, total *int64, query map[string]interface{}) error {
	if err := s.db.Model(&model.User{}).
		Where(query).Count(total).
		Scopes(pagination.Paginate(page, size)).
		Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UserRetrieve(user *model.User, id int) error {
	if err := s.db.Model(&model.User{}).
		Preload("Roles").
		First(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UserCreate(user *model.User) error {
	if err := s.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UserUpdate(user *model.User, payload map[string]interface{}) error {
	if err := s.db.Model(&user).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UserDestroy(id int) error {
	if err := s.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) UserUpdateRole(user *model.User, rolesID []int) error {
	var roles []model.Role
	if err := s.db.Model(&model.Role{}).Find(&roles, "id IN ?", rolesID).Error; err != nil {
		return err
	}
	if err := s.db.Model(&user).Association("Roles").Replace(roles); err != nil {
		return err
	}
	return nil
}
