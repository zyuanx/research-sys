package services

import (
	"context"
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type ResearchListService struct {
}

func NewResearchListService() ResearchListService {
	return ResearchListService{}
}

type IResearchListService interface {
	List(page int, size int, research *[]bson.M, total *int64) ([]*bson.M, error)
	Create(research *bson.M) error
	Update(research *bson.M, id int, data interface{}) error
	Destroy(research *bson.M, id int) error
	Retrieve(research *bson.M, id string) error
}

func (r ResearchListService) List(page int64, size int64, total *int64) ([]*bson.M, error) {

	findOptions := options.Find()
	findOptions.SetLimit(size)
	if page > 0 {
		findOptions.SetSkip(size * (page - 1))
	}

	collection := global.Mongo.Database("test").Collection("research_list")
	var err error
	*total, err = collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var cur *mongo.Cursor
	cur, err = collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	var results []*bson.M
	for cur.Next(context.TODO()) {
		var elem bson.M
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	return results, nil
}
func (r ResearchListService) Retrieve(research *bson.M, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	collection := global.Mongo.Database("test").Collection("research_list")
	if err = collection.FindOne(context.TODO(), filter).Decode(research); err != nil {
		return err
	}
	return nil
}

func (r ResearchListService) Create(research *models.ResearchList) error {
	collection := global.Mongo.Database("test").Collection("research_list")
	research.CreatedAt = time.Now()
	research.UpdatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), research)
	if err != nil {
		return err
	}
	return nil
}

func (r ResearchListService) Update(id string, data map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	data["updatedAt"] = time.Now()
	update := bson.M{
		"$set": data,
	}
	collection := global.Mongo.Database("test").Collection("research_list")
	if _, err = collection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

func (r ResearchListService) Destroy(research *bson.M, id int) error {
	panic("implement me")
}
