package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/response"
	"gin-research-sys/internal/util"
)

type IResearchService interface {
	List(researches *[]model.Research, page int, size int, total *int64, query map[string]interface{}) (err error)
	FindByAccess(researches *[]response.ResearchResponse, page int, size int,
		total *int64, query map[string]interface{}) (err error)
	Retrieve(research *model.Research, id int) error
	Create(research *model.Research) error
	Update(research *model.Research, payload map[string]interface{}) error
	Destroy(id int) error
}
type ResearchService struct{}

func NewResearchService() IResearchService {
	return ResearchService{}
}

func (r ResearchService) List(researches *[]model.Research, page int, size int,
	total *int64, query map[string]interface{}) error {
	if err := conf.Mysql.Model(&model.Research{}).
		Where(query).Count(total).
		Preload("Publisher").
		Scopes(util.Paginate(page, size)).
		Find(&researches).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) FindByAccess(researches *[]response.ResearchResponse, page int, size int,
	total *int64, query map[string]interface{}) error {
	queryAll := make(map[string]interface{})

	for key, value := range query {
		queryAll[key] = value
	}
	queryAll["access"] = "全部学院"
	if err := conf.Mysql.Model(&model.Research{}).
		Or(query).Or(queryAll).
		Count(total).Scopes(util.Paginate(page, size)).
		Find(&researches).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) Retrieve(research *model.Research, id int) error {
	if err := conf.Mysql.Model(&model.Research{}).First(&research, id).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) Create(research *model.Research) error {
	if err := conf.Mysql.Model(&model.Research{}).Create(&research).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) Update(research *model.Research, payload map[string]interface{}) error {
	if err := conf.Mysql.Model(&research).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) Destroy(id int) error {
	panic("implement me")
}
