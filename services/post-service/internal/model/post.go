package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post - структура поста,
// которая хранится в Mongo
type Post struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	AuthorID   primitive.ObjectID   `bson:"author_id"`
	Content    string               `bson:"content"`
	LikesCount int                  `bson:"likes_count"`
	Likers     []primitive.ObjectID `bson:"likers"`
	CreatedAt  time.Time            `bson:"created_at"`
}
