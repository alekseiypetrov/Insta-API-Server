package database

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo - метод, осуществляющий подключение к MongoDB
func ConnectMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)

	return client, err
}