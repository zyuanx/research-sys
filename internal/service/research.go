package service

import (
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/pagination"
)

func (s *Service) ResearchList(researches *[]model.Research, page int, size int,
	total *int64, query map[string]interface{}) error {
	if err := s.db.Model(&model.Research{}).
		Where(query).Count(total).
		Preload("Publisher").
		Scopes(pagination.Paginate(page, size)).
		Find(&researches).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ResearchRetrieve(research *model.Research, id uint64) error {
	if err := s.db.Model(&model.Research{}).First(&research, id).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {
		// 	return nil
		// }
		return err
	}
	return nil
}

func (s *Service) ResearchCreate(research *model.Research) error {
	if err := s.db.Model(&model.Research{}).Create(&research).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ResearchUpdate(research *model.Research, payload map[string]interface{}) error {
	if err := s.db.Model(&research).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ResearchDelete(id uint64) error {
	if err := s.db.Delete(&model.Research{}, id).Error; err != nil {
		return err
	}
	return nil
}
