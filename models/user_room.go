package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	CreatedAt    int    `bson:"created_at"`
	UpdatedAt    int    `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

// GetUserRoomByUserAndRoomIdentity 根据RoomID和UserID查找User是否在Room内
func GetUserRoomByUserAndRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	userRoom := new(UserRoom)
	err := Mongo.Collection(userRoom.CollectionName()).FindOne(context.Background(), bson.D{
		{"user_identity", userIdentity},
		{"room_identity", roomIdentity},
	}).Decode(userRoom)
	return userRoom, err
}

// GetUserGroupByRoomID 根据房间号取到房间内的用户
func GetUserGroupByRoomID(roomIdenty string) ([]*UserRoom, error) {
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{
		{"room_identity", roomIdenty},
	})
	if err != nil {
		return nil, err
	}
	users := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		user := new(UserRoom)
		err := cursor.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
