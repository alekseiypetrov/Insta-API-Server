package main

import (
	"user-service/internal/middleware"
	"user-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userHandler := handler.NewUserHandler()
	
	auth := r.Group("/auth")
	{
		auth.POST("/auth/login", userHandler.LoginUser)
	}

	users := r.Group("/users") 
	{
		users.POST("", userHandler.RegisterUser)

		users.GET("/me", middleware.AuthMiddleware(), userHandler.GetMe)
		users.PUT("/me/avatar", middleware.AuthMiddleware(), userHandler.UpdateAvatar)

		users.GET("/:id", userHandler.GetUser)
		users.POST("/:id/follow", middleware.AuthMiddleware(), userHandler.SetFollow)
		users.DELETE("/:id/follow", middleware.AuthMiddleware(), userHandler.DeleteFollow)
		users.GET("/:id/following", userHandler.GetFollowing)
	}
	
	r.Run(":8080")
}
