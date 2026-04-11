package middleware

import (
	"project/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	headerOfToken    = "Authorization"
	contextUserIDKey = "user_id"
	contextTagKey    = "tag"
	bearerPrefix     = "Bearer "
)

// AuthMiddleware - метод проверяет токен пользователя
func AuthMiddleware(m *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(headerOfToken)
		if authHeader == "" || len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}

		tokenString := authHeader[len(bearerPrefix):]

		resp, err := m.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Set(contextUserIDKey, resp.ID)
		c.Set(contextTagKey, resp.Tag)
		c.Next()
	}
}
