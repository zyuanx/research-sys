package main

import (
	"gin-research-sys/apps"
	"gin-research-sys/common/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	InitConfig()
	database.InitDB()
	r := gin.Default()
	r = apps.InitRouter()
	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}

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
