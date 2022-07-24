package models

import (
	"context"
	"time"
)

type MessageBasic struct {
	MessageID    string `bson:"message_identity"`
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	Data         string `bson:"data"`
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

type MessageStruct struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}

func SaveMessageToMongo(userIdentity string, msg *MessageStruct) error {
	message := MessageBasic{
		UserIdentity: userIdentity,
		RoomIdentity: msg.RoomIdentity,
		Data:         msg.Message,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	_, err := Mongo.Collection(message.CollectionName()).InsertOne(context.Background(), &message)
	return err
}
