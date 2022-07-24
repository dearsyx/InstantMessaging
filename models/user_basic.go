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

func CollectName() string {
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

func GetUserByIdentity(identity string) (*UserBasic, error) {
	user := NewUserBasic()
	err := Mongo.Collection(user.CollectionName()).FindOne(
		context.Background(), bson.D{
			{"_id", identity},
		},
	).Decode(user)
	return user, err
}

// GetUserByEmail 查看使用该邮箱的用户个数
func GetUserByEmail(email string) (int64, error) {
	return Mongo.Collection(CollectName()).CountDocuments(context.Background(), bson.D{
		{"email", email},
	})
}
