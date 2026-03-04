package middleware

import "github.com/gin-gonic/gin"

// AuthMiddleware - метод проверяет токен пользователя
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString =  "hello"
		if tokenString == "" {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}