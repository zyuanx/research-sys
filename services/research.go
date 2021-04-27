package services

import (
	"context"
	"gin-research-sys/pkg/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResearchListService struct {
}

func NewResearchListService() ResearchListService {
	return ResearchListService{}
}

type IResearchListService interface {
	Retrieve(research *bson.M, id string) error
}

func (r ResearchListService) Retrieve(research *bson.M, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	collection := global.Mongo.Database("test").Collection("research_list")
	if err := collection.FindOne(context.TODO(), filter).Decode(research); err != nil {
		return err
	}
	return nil
}
