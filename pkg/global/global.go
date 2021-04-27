package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Mysql   *gorm.DB
	Mongo *mongo.Client
)
