package app

import (
	"net/http"
	"project/pkg/jwt"
	"project/pkg/middleware"
	"project/pkg/observability"
	"project/services/post-service/internal/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func setupRoutes(r *gin.Engine, h *handler.PostHandler, m *jwt.Manager, s *observability.Stats) {
	r.Use(middleware.StatsMiddleware(s))
	r.Use(middleware.MetricsMiddleware("post-service"))
	{
		posts := r.Group("/posts")
		posts.GET("/:id", h.GetPost)

		postsWithAuth := posts.Use(middleware.AuthMiddleware(m))
		postsWithAuth.POST("", h.CreatePost)
		postsWithAuth.POST("/:id/like", h.SetLike)
		postsWithAuth.DELETE("/:id/like", h.DeleteLike)
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
		r.GET("/users/:id/posts", h.GetAllPostsOfUser)
	}
}
