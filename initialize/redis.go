package initialize

import (
	"gin-research-sys/pkg/global"
	"github.com/go-redis/redis/v8"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "play2129963", // no password set
		DB:       0,  // use default DB
	})
	global.Redis = rdb
}
