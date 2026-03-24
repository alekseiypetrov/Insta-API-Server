package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User - модель пользователя
type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Tag            string               `bson:"tag"`
	Password       string               `bson:"password"`
	SecretAnswer   string               `bson:"secret_answer"`
	AvatarURL      string               `bson:"avatar_url"`
	FollowingCount int                  `bson:"following_count"`
	FollowersCount int                  `bson:"followers_count"`
	Followers      []primitive.ObjectID `bson:"followers"`
	Following      []primitive.ObjectID `bson:"following"`
}

// LoginUser - модель пользователя для авторизации
type LoginUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Tag      string             `bson:"tag"`
	Password string             `bson:"password"`
}
