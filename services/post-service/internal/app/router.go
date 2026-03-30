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
		posts.GET("/:id", h.GetPost)

		postsWithAuth := posts.Use(middleware.AuthMiddleware(m))
		postsWithAuth.POST("", h.CreatePost)
		postsWithAuth.POST("/:id/like", h.SetLike)
		postsWithAuth.DELETE("/:id/like", h.DeleteLike)
	}
	{
		r.GET("/users/:id/posts", h.GetAllPostsOfUser)
	}
}
