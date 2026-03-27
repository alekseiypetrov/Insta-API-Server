package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config - стурктура,
// хранящая значения по ключам env-файла
type Config struct {
	JWTSecret   string
}

// LoadConfig - конструктор конфига
// при отсутствии одного из ключа возвращает ошибку
func LoadConfig() (*Config, error) {
	godotenv.Load()

	cfg := &Config{
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}

	switch {
	case cfg.JWTSecret == "":
		return nil, fmt.Errorf("JWT_SECRET is not set")
	default:
		return cfg, nil
	}
}
