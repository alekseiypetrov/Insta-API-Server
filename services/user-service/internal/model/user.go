package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	// User - модель пользователя
	User struct {
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
	LoginUser struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Tag      string             `bson:"tag"`
		Password string             `bson:"password"`
	}

	// FollowingResult - модель списка подписок пользователя
	FollowingResult struct {
		Following []primitive.ObjectID `bson:"following"`
	}
)
