package service

import (
	"fmt"
	"project/services/user-service/internal/dto"
	"project/services/user-service/internal/repository"
	"strings"

	"project/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

// AuthService - сервис, выполняющий
// бизнес-логику регистрации и авторизации
type AuthService struct {
	userRepo   *repository.UserRepository
	jwtManager *jwt.Manager
}

// NewAuthService - конструктор сервиса
func NewAuthService(r *repository.UserRepository, m *jwt.Manager) *AuthService {
	return &AuthService{
		userRepo:   r,
		jwtManager: m,
	}
}

// SignUp - регистрация нового пользователя
func (s *AuthService) SignUp(req dto.RegisterRequest) (string, error) {
	if _, ok := s.userRepo.QueryUser(strings.ToLower(req.Tag)); ok {
		return "", fmt.Errorf("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(req.SecretAnswer), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	id, err := s.userRepo.CreateUser(req.Tag, string(hashedPassword), string(hashedAnswer))
	if err != nil {
		return "", fmt.Errorf("failed to create user")
	}

	token, err := s.jwtManager.GenerateToken(id, req.Tag)
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	return token, nil
}

// SignIn - вход в аккаунт
func (s *AuthService) SignIn(req dto.LoginRequest) (string, error) {
	user, ok := s.userRepo.QueryUser(strings.ToLower(req.Tag))
	if !ok {
		return "", fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := s.jwtManager.GenerateToken(user.ID.Hex(), req.Tag)
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	return token, nil
}
