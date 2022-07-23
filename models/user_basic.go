package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `bson:"_id"`
	Account   string `bson:"account"`
	Username  string `bson:"username"`
	Password  string `bson:"password"`
	Avatar    string `bson:"avatar"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (*UserBasic) CollectionName() string {
	return "user_basic"
}

func NewUserBasic() *UserBasic {
	return &UserBasic{}
}

func GetUserByAccountPassword(account, password string) (*UserBasic, error) {
	user := NewUserBasic()
	err := Mongo.Collection(user.CollectionName()).FindOne(
		context.Background(), bson.D{
			{"account", account},
			{"password", password},
		},
	).Decode(user)
	return user, err
}
