package initialize

import (
	"fmt"
	"gin-research-sys/internal/conf"
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
	conf.Redis = rdb
}
