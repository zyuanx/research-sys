package services

import (
	"context"
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResearchService struct{}

func NewResearchService() IResearchService {
	return ResearchService{}
}

type IResearchService interface {
	List(page int, size int, researches *[]models.Research, total *int64) error
	Retrieve(research *models.Research, id int) error
	Create(research *models.Research) error
	Update(research *models.Research) error
	Destroy(research *bson.M, id int) error
}

func (r ResearchService) List(page int, size int, researches *[]models.Research, total *int64) error {
	if err := global.Mysql.Model(&models.Research{}).
		Count(total).
		Preload("User").
		Scopes(global.Paginate(page, size)).
		Find(&researches).Error; err != nil {
		return err
	}
	return nil
	//findOptions := options.Find()
	//findOptions.SetLimit(size)
	//if page > 0 {
	//	findOptions.SetSkip(size * (page - 1))
	//}
	//
	//collection := global.Mongo.Database("test").Collection("research_list")
	//var err error
	//*total, err = collection.CountDocuments(context.TODO(), bson.M{})
	//if err != nil {
	//	return nil, err
	//}
	//var cur *mongo.Cursor
	//cur, err = collection.Find(context.TODO(), bson.M{}, findOptions)
	//if err != nil {
	//	return nil, err
	//}
	//var results []*bson.M
	//for cur.Next(context.TODO()) {
	//	var elem bson.M
	//	err = cur.Decode(&elem)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	results = append(results, &elem)
	//}
	//return results, nil
}

func (r ResearchService) Retrieve(research *models.Research, id int) error {
	if err := global.Mysql.Model(&models.Research{}).First(&research, id).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) Create(research *models.Research) error {
	if err := global.Mysql.Model(&models.Research{}).Create(&research).Error; err != nil {
		return err
	}
	return nil
}

func (r ResearchService) Update(research *models.Research) error {

	if err := global.Mysql.Save(&research).Error; err != nil {
		return err
	}
	return nil
	//objectID, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return err
	//}
	//filter := bson.M{"_id": objectID}
	//data["updatedAt"] = time.Now()
	//update := bson.M{
	//	"$set": data,
	//}
	//collection := global.Mongo.Database("test").Collection("research_list")
	//if _, err = collection.UpdateOne(context.TODO(), filter, update); err != nil {
	//	return err
	//}
	//return nil
}

func (r ResearchService) Destroy(research *bson.M, id int) error {
	panic("implement me")
}

type ResearchMgoService struct{}

func NewResearchMgoService() IResearchMgoService {
	return ResearchMgoService{}
}

type IResearchMgoService interface {
	Create(research *models.ResearchMgo) (*mongo.InsertOneResult, error)
	Retrieve(research *models.ResearchMgo, id string) error
}

func (r ResearchMgoService) Create(research *models.ResearchMgo) (*mongo.InsertOneResult, error) {
	collection := global.Mongo.
		Database("test").
		Collection("research")
	one, err := collection.InsertOne(context.TODO(), research)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (r ResearchMgoService) Retrieve(research *models.ResearchMgo, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	collection := global.Mongo.Database("test").Collection("research")
	if err = collection.FindOne(context.TODO(), filter).Decode(&research); err != nil {
		return err
	}
	return nil
}
