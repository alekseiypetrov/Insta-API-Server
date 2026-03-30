package app

import (
	"project/pkg/jwt"
	"project/pkg/middleware"
	"project/services/feed-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine, h *handler.FeedHandler, m *jwt.Manager) {
	{
		feed := r.Group("/feed")
		feed.GET("/me", middleware.AuthMiddleware(m), h.GetFeed)
	}
}
