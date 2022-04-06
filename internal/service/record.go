package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/util"
)

type RecordService struct{}

func NewRecordService() IRecordService {
	return RecordService{}
}

type IRecordService interface {
	List(page int, size int, records *[]model.Record, total *int64, query map[string]interface{}) error
	Retrieve(record *model.Record, id int) error
	Create(record *model.Record) error

	FindByResearchID(records *model.Record, researchId string, id int) error
	ListID(id string, records *[]model.Record, total *int64) error
}

func (r RecordService) List(page int, size int, records *[]model.Record, total *int64, query map[string]interface{}) error {
	if err := conf.Mysql.Model(&model.Record{}).
		Where(query).
		Count(total).
		Scopes(util.Paginate(page, size)).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}

func (r RecordService) Retrieve(record *model.Record, id int) error {
	if err := conf.Mysql.Model(&model.Record{}).First(&record, id).Error; err != nil {
		return err
	}
	return nil
}

func (r RecordService) Create(record *model.Record) error {
	if err := conf.Mysql.Model(&model.Record{}).Create(&record).Error; err != nil {
		return err
	}
	return nil
}

func (r RecordService) FindByResearchID(records *model.Record, researchId string, id int) error {
	if err := conf.Mysql.Model(&model.Record{}).
		Where("research_id = ?", researchId).
		Where("user_id = ?", id).
		First(&records).Error; err != nil {
		return err
	}
	return nil
}
func (r RecordService) ListID(id string, records *[]model.Record, total *int64) error {
	if err := conf.Mysql.Model(&model.Record{}).
		Count(total).
		Preload("User").
		Where("research_id = ?", id).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}

type OpenRecordService struct{}

func NewOpenRecordService() IOpenRecordService {
	return OpenRecordService{}
}

type IOpenRecordService interface {
	List(page int, size int, records *[]model.OpenRecord, total *int64, query map[string]interface{}) error
	Retrieve(record *model.OpenRecord, id int) error
	Create(record *model.OpenRecord) error
	ListByResearchID(records *[]model.OpenRecord, researchID uint) error
}

func (o OpenRecordService) List(page int, size int, records *[]model.OpenRecord,
	total *int64, query map[string]interface{}) error {
	if err := conf.Mysql.Model(&model.OpenRecord{}).
		Where(query).
		Count(total).
		Scopes(util.Paginate(page, size)).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}

func (o OpenRecordService) Retrieve(record *model.OpenRecord, id int) error {
	if err := conf.Mysql.Model(&model.OpenRecord{}).First(&record, id).Error; err != nil {
		return err
	}
	return nil
}

func (o OpenRecordService) Create(record *model.OpenRecord) error {
	if err := conf.Mysql.Model(&model.OpenRecord{}).Create(&record).Error; err != nil {
		return err
	}
	return nil
}

func (o OpenRecordService) ListByResearchID(records *[]model.OpenRecord, researchID uint) error {
	if err := conf.Mysql.Model(&model.OpenRecord{}).
		Where("research_id = ?", researchID).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}
