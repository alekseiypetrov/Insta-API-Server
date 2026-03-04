package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User - модель пользователя
type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Tag string `bson:"tag" json:"tag"`
	Password string `bson:"password"`
	SecretAnswer string `bson:"secret_answer"`
	AvatarURL string `bson:"avatar_url" json:"avatar_url"`
	Following []primitive.ObjectID `bson:"following" json:"following"`
	FollowersCount int `bson:"followers_count" json:"followers_count"`
}