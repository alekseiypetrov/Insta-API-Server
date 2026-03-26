package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config - стурктура,
// хранящая значения по ключам env-файла
type Config struct {
	MongoURI            string
	MongoDBName         string
	MongoCollectionName string
	JWTSecret           string
}

// LoadConfig - конструктор конфига
// при отсутствии одного из ключа возвращает ошибку
func LoadConfig() (*Config, error) {
	godotenv.Load()

	cfg := &Config{
		MongoURI:            os.Getenv("MONGO_URI"),
		MongoDBName:         os.Getenv("MONGO_DB"),
		MongoCollectionName: os.Getenv("MONGO_COLLECTION_NAME"),
		JWTSecret:           os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set")
	} else if cfg.MongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	} else if cfg.MongoDBName == "" {
		return nil, fmt.Errorf("MONGO_DB is not set")
	} else if cfg.MongoCollectionName == "" {
		return nil, fmt.Errorf("MONGO_COLLECTION_NAME is not set")
	}

	return cfg, nil
}
