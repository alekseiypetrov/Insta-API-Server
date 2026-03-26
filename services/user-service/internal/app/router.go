package app

import (
	"project/services/user-service/internal/handler"

	"project/pkg/jwt"
	"project/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine, h *handler.UserHandler, m *jwt.Manager) {
	{
		auth := r.Group("/auth")
		auth.POST("/login", h.LoginUser)
		auth.POST("/register", h.RegisterUser)
	}

	{
		users := r.Group("/users")
		users.GET("/:id", h.GetUser)

		authUsers := users.Group("")
		authUsers.Use(middleware.AuthMiddleware(m))
		authUsers.GET("/me", h.GetMe)
		// authUsers.PUT("/me/avatar", h.UpdateAvatar)
		authUsers.POST("/:id/follow", h.SetFollow)
		authUsers.DELETE("/:id/follow", h.DeleteFollow)
	}
}
