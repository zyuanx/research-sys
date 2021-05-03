package global

import (
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Mysql       *gorm.DB
	Mongo       *mongo.Client
	Enforcer    *casbin.Enforcer
	sugarLogger *zap.SugaredLogger
)
