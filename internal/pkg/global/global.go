package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	RootDir string
	MySQL   *gorm.DB
	Redis   *redis.Client
)
