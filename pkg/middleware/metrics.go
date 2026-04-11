package middleware

import (
	"strconv"

	"project/pkg/observability"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		observability.HttpRequests.WithLabelValues(serviceName, status).Inc()
	}
}
