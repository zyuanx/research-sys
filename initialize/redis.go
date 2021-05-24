package initialize

import (
	"fmt"
	"gin-research-sys/pkg/global"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func Redis() {
	addr := fmt.Sprintf("%s:%d",
		viper.GetString("redis.host"),
		viper.GetInt("redis.port"),
	)
	password := viper.GetString("redis.password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	global.Redis = rdb
}
