package app

import (
	"project/pkg/jwt"
	"project/pkg/middleware"
	"project/services/post-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine, h *handler.PostHandler, m *jwt.Manager) {
	{
		posts := r.Group("/posts")
		posts.POST("", middleware.AuthMiddleware(m), h.CreatePost)
		posts.GET("/:id", h.GetPost)
		posts.POST("/:id/like", middleware.AuthMiddleware(m), h.SetLike)
		posts.DELETE("/:id/like", middleware.AuthMiddleware(m), h.DeleteLike)
	}
	{
		r.GET("/users/:id/posts", h.GetAllPostsOfUser)
	}
}
