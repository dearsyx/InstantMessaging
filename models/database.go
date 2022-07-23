package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo *mongo.Database

func MongoInit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		SetAuth(options.Credential{
			Username: "admin",
			Password: "123456",
		}).
		ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalln(err)
	}
	Mongo = client.Database("instant_messaging")
}
