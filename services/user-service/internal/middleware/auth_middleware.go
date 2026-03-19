package middleware

import (
	"user-service/internal/jwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware - метод проверяет токен пользователя
func AuthMiddleware(m *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(401)
		}

		user_id, error := m.VerifyToken(tokenString) //jwt.ParseToken(tokenString)
		if error != nil {
			c.AbortWithStatus(401)
			return
		}
		c.AddParam("user_id", string(user_id))
		c.Next()
	}
}
