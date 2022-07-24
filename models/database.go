package models

import (
	"code.project.com/InstantMessaging/pkg/config"
	"context"
	"fmt"
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
			Username: config.Config.Mongo.Username,
			Password: config.Config.Mongo.Password,
		}).ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.Config.Mongo.Host, config.Config.Mongo.Port)))
	if err != nil {
		log.Fatalln(err)
	}
	Mongo = client.Database("instant_messaging")
}
