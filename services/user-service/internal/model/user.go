package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User - модель пользователя
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Tag            string             `bson:"tag" json:"tag"`
	Password       string             `bson:"password"`
	SecretAnswer   string             `bson:"secret_answer"`
	AvatarURL      string             `bson:"avatar_url" json:"avatar_url"`
	FollowingCount      int                `bson:"following_count" json:"following_count"`
	FollowersCount int                `bson:"followers_count" json:"followers_count"`
}

// LoginUser - модель пользователя для авторизации
type LoginUser struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Tag            string             `bson:"tag"`
	Password       string             `bson:"password"`
}