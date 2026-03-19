package app

import (
	"user-service/internal/config"
	"user-service/internal/database"
	"user-service/internal/handler"
	"user-service/internal/jwt"
	"user-service/internal/repository"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

// App (структура) - экземпляр приложения (микросервиса)
type App struct {
	Router *gin.Engine
}

// NewApp - конструктор экземпляра приложения
func NewApp() (*App, error) {
	envConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	db, err := database.ConnectMongo(envConfig.MongoURI)
	if err != nil {
		return nil, err
	}
	jwtManager := jwt.NewManager(envConfig.JWTSecret)
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, jwtManager)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(authService, userService)

	setupRoutes(r, userHandler, jwtManager)

	return &App{Router: r}, nil
}
