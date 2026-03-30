package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config - стурктура,
// хранящая значения по ключам env-файла
type Config struct {
	JWTSecret      string
	UserServiceURL string
	PostServiceURL string
}

// LoadConfig - конструктор конфига
// при отсутствии одного из ключа возвращает ошибку
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		JWTSecret:      os.Getenv("JWT_SECRET"),
		UserServiceURL: os.Getenv("USER_SERVICE_URL"),
		PostServiceURL: os.Getenv("POST_SERVICE_URL"),
	}

	switch {
	case cfg.JWTSecret == "":
		return nil, fmt.Errorf("JWT_SECRET is not set")
	case cfg.UserServiceURL == "":
		return nil, fmt.Errorf("USER_SERVICE_URL is not set")
	case cfg.PostServiceURL == "":
		return nil, fmt.Errorf("POST_SERVICE_URL is not set")
	default:
		return cfg, nil
	}
}
