package global

import (
	"github.com/allegro/bigcache/v3"
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Mysql    *gorm.DB
	Mongo    *mongo.Client
	Cache    *bigcache.BigCache
	Enforcer *casbin.Enforcer
)
