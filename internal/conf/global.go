package conf

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Mysql    *gorm.DB
	Enforcer *casbin.Enforcer
	Redis    *redis.Client
)
