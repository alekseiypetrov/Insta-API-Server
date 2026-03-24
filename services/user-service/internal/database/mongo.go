package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo - структура, которая хранит клиент и БД
type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// ConnectMongo - подключение к MongoDB
func ConnectMongo(uri, dbname string) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping failed: %w", err)
	}

	db := client.Database(dbname)

	return &Mongo{
		Client: client,
		DB:     db,
	}, nil
}
