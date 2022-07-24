package test

import (
	"code.project.com/InstantMessaging/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongoFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		SetAuth(options.Credential{
			Username: "admin",
			Password: "123456",
		}).
		ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("instant_messaging")
	user := new(models.UserBasic)
	err = db.Collection(user.CollectionName()).FindOne(context.Background(), bson.D{}).Decode(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*user)
}

func TestMongoFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		SetAuth(options.Credential{
			Username: "admin",
			Password: "123456",
		}).
		ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("instant_messaging")
	cursor, err := db.Collection("user_room").Find(context.Background(), bson.D{
		{"room_identity", "room_identity"},
	})
	if err != nil {
		t.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		ub := new(models.UserRoom)
		err := cursor.Decode(ub)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(*ub)
	}
}
