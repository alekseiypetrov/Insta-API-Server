package app

import (
	"project/services/user-service/internal/config"
	"project/services/user-service/internal/handler"
	"project/services/user-service/internal/repository"
	"project/services/user-service/internal/service"

	"project/pkg/database"
	"project/pkg/jwt"

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

	mongo, err := database.ConnectMongo(envConfig.MongoURI, envConfig.MongoDBName)
	if err != nil {
		return nil, err
	}
	jwtManager := jwt.NewManager(envConfig.JWTSecret)

	userRepo := repository.NewUserRepository(mongo)
	if err := userRepo.CreateIndexes(); err != nil {
		return nil, err
	}

	authService := service.NewAuthService(userRepo, jwtManager)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(authService, userService)

	setupRoutes(r, userHandler, jwtManager)

	return &App{Router: r}, nil
}
