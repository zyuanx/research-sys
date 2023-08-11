package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	RootDir string
	MySQL   *gorm.DB
	// Mysql   *gorm.DB
	Redis *redis.Client
)
