package models

type UserBasic struct {
	Identity  string `bson:"_id"`
	UserID    string `bson:"user_id"`
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
