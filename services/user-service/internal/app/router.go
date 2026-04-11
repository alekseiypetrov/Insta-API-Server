package app

import (
	"net/http"
	"project/services/user-service/internal/handler"
	"time"

	"project/pkg/jwt"
	"project/pkg/middleware"
	"project/pkg/observability"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func setupRoutes(r *gin.Engine, h *handler.UserHandler, m *jwt.Manager, s *observability.Stats) {
	r.Use(middleware.StatsMiddleware(s))
	r.Use(middleware.MetricsMiddleware("user-service"))

	{
		auth := r.Group("/auth")
		auth.POST("/login", h.LoginUser)
		auth.POST("/register", h.RegisterUser)
	}
	{
		r.GET("/stats", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version":        s.Version,
				"started_at":     s.StartedAt,
				"uptime":         time.Since(s.StartedAt).String(),
				"requests_total": s.RequestsTotal,
				"responses":      s.Responses,
			})
		})
		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
	{
		users := r.Group("/users")
		users.GET("/:id", h.GetUser)

		authUsers := users.Group("")
		authUsers.Use(middleware.AuthMiddleware(m))
		authUsers.GET("/me", h.GetMe)
		authUsers.GET("/me/following", h.GetFollowings)
		// authUsers.PUT("/me/avatar", h.UpdateAvatar)
		authUsers.POST("/:id/follow", h.SetFollow)
		authUsers.DELETE("/:id/follow", h.DeleteFollow)
	}
}
