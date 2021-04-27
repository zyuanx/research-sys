package initialize

import (
	"context"
	"gin-research-sys/pkg/global"
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
	uri := "mongodb://admin:123456@127.0.0.1:27017/?authSource=admin"
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
