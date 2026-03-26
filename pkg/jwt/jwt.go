package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Manager - менеджер по JWT-токенам
// он их создает, выдает и проверяет
type Manager struct {
	secret string
}

// NewManager - конструктор менеджера JWT
func NewManager(secret string) *Manager {
	return &Manager{
		secret: secret,
	}
}

// GenerateToken - создает JWT-токен
func (m *Manager) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(48 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

// VerifyToken - разборка и проверка токена
func (m *Manager) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(m.secret), nil
	})

	switch {
	case err != nil:
		return "", err
	case !token.Valid:
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("user_id not found")
	}

	return userID, nil
}
