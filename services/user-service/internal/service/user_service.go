package service

import (
	"fmt"
	"user-service/internal/model"
	"user-service/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserService - сервис, выполняющий
// бизнес-логику работы с профилем и
// взаимодействием с другими пользователями
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService - конструктор сервиса
func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: r,
	}
}

// GetByID - метод, возвращающий информацию
// о пользователе по id
func (s *UserService) GetByID(userID string) (model.User, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return model.User{}, fmt.Errorf("invalid user id")
	}

	user, err := s.userRepo.QueryProfile(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

// TODO: - Will be done later

// UpdateAvatar - метод, обновляющий аватар пользователя
func (s *UserService) UpdateAvatar() {}

// FollowUser - метод, оформляющий подписку на пользователя
func (s *UserService) FollowUser() {}

// UnfollowUser - метод, обратный FollowUser
func (s *UserService) UnfollowUser() {}

// GetFollowers - метод, возвращающий подписчиков
func (s *UserService) GetFollowers() {}
