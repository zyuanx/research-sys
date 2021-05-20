package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Mysql    *gorm.DB
	Mongo    *mongo.Client
	Enforcer *casbin.Enforcer
	Redis    *redis.Client
)
