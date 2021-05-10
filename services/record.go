package services

import (
	"context"
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordService struct{}

func NewRecordService() IRecordService {
	return RecordService{}
}

type IRecordService interface {
	List(page int, size int, records *[]models.Record, total *int64) error
	Retrieve(record *models.Record, id int) error
	Create(record *models.Record) error

	ListID(id string, records *[]models.Record, total *int64) error
}

func (r RecordService) List(page int, size int, records *[]models.Record, total *int64) error {
	if err := global.Mysql.Model(&models.Record{}).
		Count(total).
		Preload("User").
		Scopes(global.Paginate(page, size)).
		Find(&records).Error; err != nil {
		return err
	}
	return nil
}

func (r RecordService) Retrieve(record *models.Record, id int) error {
	if err := global.Mysql.Model(&models.Record{}).First(&record, id).Error; err != nil {
		return err
	}
	return nil
}

func (r RecordService) Create(record *models.Record) error {
	if err := global.Mysql.Model(&models.Record{}).Create(&record).Error; err != nil {
		return err
	}
	return nil
}

func (r RecordService) ListID(id string, records *[]models.Record, total *int64) error {
	if err := global.Mysql.Model(&models.Record{}).
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
	Create(research *models.RecordMgo) (*mongo.InsertOneResult, error)
	Retrieve(research *models.RecordMgo, id string) error
}

func (r RecordMgoService) Create(record *models.RecordMgo) (*mongo.InsertOneResult, error) {
	collection := global.Mongo.
		Database("test").
		Collection("record")

	one, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (r RecordMgoService) Retrieve(research *models.RecordMgo, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	collection := global.Mongo.
		Database("test").
		Collection("record")
	if err = collection.FindOne(context.TODO(), filter).Decode(&research); err != nil {
		return err
	}
	return nil
}
