package service

import (
	"gin-research-sys/internal/conf"
	"gin-research-sys/internal/model"
	"gin-research-sys/internal/response"
	"gin-research-sys/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

type ResearchService struct{}

func NewResearchService() IResearchService {
	return &ResearchService{}
}

type IResearchService interface {
	List(page int, size int, researches *[]model.Research,
		total *int64, query map[string]interface{}) (err error)
	FindByAccess(page int, size int, researches *[]response.ResearchResponse,
		total *int64, query map[string]interface{}) (err error)
	Retrieve(research *model.Research, id int) error
	Create(research *model.Research) error
	Update(research *model.Research) error
	Destroy(research *bson.M, id int) error
}

func (r *ResearchService) List(page int, size int, researches *[]model.Research,
	total *int64, query map[string]interface{}) (err error) {
	err = conf.Mysql.Model(&model.Research{}).
		Where(query).
		Count(total).
		Preload("User").
		Scopes(util.Paginate(page, size)).
		Find(&researches).Error
	return
}

func (r *ResearchService) FindByAccess(page int, size int, researches *[]response.ResearchResponse,
	total *int64, query map[string]interface{}) (err error) {
	queryAll := make(map[string]interface{})

	for key, value := range query {
		queryAll[key] = value
	}
	queryAll["access"] = "全部学院"
	err = conf.Mysql.Model(&model.Research{}).
		Or(query).Or(queryAll).
		Count(total).Scopes(util.Paginate(page, size)).
		Find(&researches).Error
	return
}

func (r *ResearchService) Retrieve(research *model.Research, id int) error {
	if err := conf.Mysql.Model(&model.Research{}).First(&research, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ResearchService) Create(research *model.Research) error {
	if err := conf.Mysql.Model(&model.Research{}).Create(&research).Error; err != nil {
		return err
	}
	return nil
}

func (r *ResearchService) Update(research *model.Research) error {

	if err := conf.Mysql.Save(&research).Error; err != nil {
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
	//collection := conf.Mongo.Database("test").Collection("research_list")
	//if _, err = collection.UpdateOne(context.TODO(), filter, update); err != nil {
	//	return err
	//}
	//return nil
}

func (r *ResearchService) Destroy(research *bson.M, id int) error {
	panic("implement me")
}

//type ResearchMgoService struct{}
//
//func NewResearchMgoService() IResearchMgoService {
//	return ResearchMgoService{}
//}
//
//type IResearchMgoService interface {
//	Create(research *model.ResearchMgo) (*mongo.InsertOneResult, error)
//	Retrieve(research *model.ResearchMgo, id string) error
//}
//
//func (r ResearchMgoService) Create(research *model.ResearchMgo) (*mongo.InsertOneResult, error) {
//	collection := conf.Mongo.
//		Database("test").
//		Collection("research")
//	one, err := collection.InsertOne(context.TODO(), research)
//	if err != nil {
//		return nil, err
//	}
//	return one, nil
//}
//
//func (r ResearchMgoService) Retrieve(research *model.ResearchMgo, id string) error {
//	objectID, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return err
//	}
//	filter := bson.M{"_id": objectID}
//	collection := conf.Mongo.Database("test").Collection("research")
//	if err = collection.FindOne(context.TODO(), filter).Decode(&research); err != nil {
//		return err
//	}
//	return nil
//}
