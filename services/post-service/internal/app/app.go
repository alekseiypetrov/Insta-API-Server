package app

import (
	"project/pkg/database"
	"project/pkg/jwt"
	"project/pkg/observability"
	"project/services/post-service/internal/config"
	"project/services/post-service/internal/handler"
	"project/services/post-service/internal/repository"
	"project/services/post-service/internal/service"

	"github.com/gin-gonic/gin"
)

// App - экземпляр приложения
type App struct {
	Router *gin.Engine
}

// NewApp - конструктор экземпляра приложения
func NewApp() (*App, error) {
	r := gin.Default()
	envCfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	mongo, err := database.ConnectMongo(envCfg.MongoURI, envCfg.MongoDBName)
	if err != nil {
		return nil, err
	}

	stats := observability.NewStats("1.0.0")
	observability.InitPrometheus()
	observability.InitTracer("post-service")

	jwtManager := jwt.NewManager(envCfg.JWTSecret)
	postRepository := repository.NewPostRepository(mongo, envCfg.MongoCollectionName)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	setupRoutes(r, postHandler, jwtManager, stats)

	return &App{r}, nil
}