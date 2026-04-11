package middleware

import (
	"project/pkg/observability"

	"github.com/gin-gonic/gin"
)

func StatsMiddleware(stats *observability.Stats) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		stats.Inc(status)
	}
}