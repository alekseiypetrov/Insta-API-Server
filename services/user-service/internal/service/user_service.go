package service

import (
	"fmt"
	"project/services/user-service/internal/model"
	"project/services/user-service/internal/repository"

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

// FollowUser - метод, оформляющий подписку на пользователя
func (s *UserService) FollowUser(followerID, targetID string) error {
	if followerID == targetID {
		return fmt.Errorf("cannot follow yourself")
	}

	myID, err := primitive.ObjectIDFromHex(followerID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}
	otherID, err := primitive.ObjectIDFromHex(targetID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	if err := s.userRepo.InsertFollow(myID, otherID); err != nil {
		return err
	}

	return nil
}

// UnfollowUser - метод, обратный FollowUser
func (s *UserService) UnfollowUser(firstID, secondID string) error {
	if firstID == secondID {
		return fmt.Errorf("cannot unfollow yourself")
	}

	myID, err := primitive.ObjectIDFromHex(firstID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}
	otherID, err := primitive.ObjectIDFromHex(secondID)
	if err != nil {
		return fmt.Errorf("invalid user id")
	}

	if err := s.userRepo.DeleteFollow(myID, otherID); err != nil {
		return err
	}

	return nil
}

// TODO: - Will be done later

// UpdateAvatar - метод, обновляющий аватар пользователя
// func (s *UserService) UpdateAvatar() {}
