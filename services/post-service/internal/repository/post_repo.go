package repository

import (
	"project/pkg/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// PostRepository - структура, 
// которая выполняет запросы к БД
type PostRepository struct {
	collection *mongo.Collection
}

// NewPostRepository - конструктор репозитория
func NewPostRepository(connection *database.Mongo, collectionName string) *PostRepository {
	return &PostRepository{
		collection: connection.DB.Collection(collectionName),
	}
}
