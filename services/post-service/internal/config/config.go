package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config - стурктура,
// хранящая значения по ключам env-файла
type Config struct {
	JWTSecret           string
	MongoURI            string
	MongoDBName         string
	MongoCollectionName string
}

// LoadConfig - конструктор конфига
// при отсутствии одного из ключа возвращает ошибку
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		JWTSecret:           os.Getenv("JWT_SECRET"),
		MongoURI:            os.Getenv("MONGO_URI"),
		MongoDBName:         os.Getenv("MONGO_DB"),
		MongoCollectionName: os.Getenv("MONGO_COLLECTION_NAME"),
	}

	switch {
	case cfg.JWTSecret == "":
		return nil, fmt.Errorf("JWT_SECRET is not set")
	case cfg.MongoURI == "":
		return nil, fmt.Errorf("MONGO_URI is not set")
	case cfg.MongoDBName == "":
		return nil, fmt.Errorf("MONGO_DB is not set")
	case cfg.MongoCollectionName == "":
		return nil, fmt.Errorf("MONGO_COLLECTION_NAME is not set")
	default:
		return cfg, nil
	}
}
