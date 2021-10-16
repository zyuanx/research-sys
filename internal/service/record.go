package service

import (
	"context"
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		Preload("User").
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

type RecordMgoService struct{}

func NewRecordMgoService() IRecordMgoService {
	return RecordMgoService{}
}

type IRecordMgoService interface {
	Create(research *model.RecordMgo) (*mongo.InsertOneResult, error)
	Retrieve(research *model.RecordMgo, id string) error
}

func (r RecordMgoService) Create(record *model.RecordMgo) (*mongo.InsertOneResult, error) {
	collection := conf.Mongo.
		Database("test").
		Collection("record")

	one, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (r RecordMgoService) Retrieve(research *model.RecordMgo, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	collection := conf.Mongo.
		Database("test").
		Collection("record")
	if err = collection.FindOne(context.TODO(), filter).Decode(&research); err != nil {
		return err
	}
	return nil
}
