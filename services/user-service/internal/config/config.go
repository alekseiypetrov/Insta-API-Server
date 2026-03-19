package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config - стурктура,
// хранящая значения по ключам env-файла
type Config struct {
	MongoURI  string
	JWTSecret string
}

// LoadConfig - конструктор конфига
// при отсутствии одного из ключа возвращает ошибку
func LoadConfig() (*Config, error) {
	godotenv.Load()

	cfg := &Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set")
	} else if cfg.MongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	}

	return cfg, nil
}
