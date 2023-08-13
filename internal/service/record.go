package service

import (
	"errors"

	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/pagination"
	"gorm.io/gorm"
)

func (s *Service) RecordList(page int, size int, records *[]model.Record, total *int64, query map[string]interface{}) error {
	if err := s.db.Model(&model.Record{}).
		Where(query).
		Count(total).
		Scopes(pagination.Paginate(page, size)).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RecordRetrieve(record *model.Record, id int) error {
	if err := s.db.Model(&model.Record{}).First(&record, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) RecordCreate(record *model.Record, research *model.Research) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	rec := model.Record{}
	if err := tx.Set("gorm:query_option", "FOR UPDATE").
		Where("research_id = ?", research.ID).
		Where("user_id = ?", record.UserID).
		First(&rec).Error; err != nil {
		// 如果不是没有找到记录，则返回错误
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
	}
	// 此时指定 id 的记录被锁住。如果表中无符合记录的数据，则排他锁不生效
	// 执行其他数据库操作
	if err := tx.Create(&record).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (s *Service) RecordFindByResearchID(record *model.Record, researchId string, id int) error {
	if err := s.db.Where("research_id = ?", researchId).
		Where("user_id = ?", id).
		First(&record).Error; err != nil {
		return err
	}
	return nil
}
func (s *Service) RecordListID(id string, records *[]model.Record, total *int64) error {
	if err := s.db.Model(&model.Record{}).
		Count(total).
		Preload("User").
		Where("research_id = ?", id).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}
