package app

import (
	"net/http"
	"project/pkg/jwt"
	"project/pkg/middleware"
	"project/pkg/observability"
	"project/services/feed-service/internal/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func setupRoutes(r *gin.Engine, h *handler.FeedHandler, m *jwt.Manager, s *observability.Stats) {
	r.Use(middleware.StatsMiddleware(s))
	r.Use(middleware.MetricsMiddleware("feed-service"))
	{
		feed := r.Group("/feed")
		feed.GET("/me", middleware.AuthMiddleware(m), h.GetFeed)
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
}
