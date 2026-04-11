package app

import (
	"project/pkg/jwt"
	"project/pkg/observability"
	"project/services/feed-service/internal/client"
	"project/services/feed-service/internal/config"
	"project/services/feed-service/internal/handler"
	"project/services/feed-service/internal/service"

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

	stats := observability.NewStats("1.0.0")
	observability.InitPrometheus()
	observability.InitTracer("feed-service")

	userClient := client.NewUserClient(envCfg.UserServiceURL)
	postClient := client.NewPostClient(envCfg.PostServiceURL)
	jwtManager := jwt.NewManager(envCfg.JWTSecret)
	feedService := service.NewFeedService(userClient, postClient)
	feedHandler := handler.NewFeedHandler(feedService)

	setupRoutes(r, feedHandler, jwtManager, stats)

	return &App{r}, nil
}
