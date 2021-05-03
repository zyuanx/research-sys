package initialize

import (
	"context"
	"fmt"
	"gin-research-sys/pkg/global"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func MongoDB() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=%s",
		viper.GetString("mongodb.username"),
		viper.GetString("mongodb.password"),
		viper.GetString("mongodb.host"),
		viper.GetInt("mongodb.port"),
		viper.GetString("mongodb.authSource"),
	)
	//uri := "mongodb://admin:123456@127.0.0.1:27017/?authSource=admin"
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("DB %v err: %v", uri, err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	global.Mongo = client
}
