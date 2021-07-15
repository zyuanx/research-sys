package initialize

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic("read config error")
	}

}
