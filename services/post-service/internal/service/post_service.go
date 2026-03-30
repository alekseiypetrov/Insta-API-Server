package service

import (
	"fmt"
	"project/services/post-service/internal/dto"
	"project/services/post-service/internal/model"
	"project/services/post-service/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostService - структура,
// выполняющая бизнес-логику микросервиса
type PostService struct {
	repository *repository.PostRepository
}

// NewPostService - конструктор PostService
func NewPostService(r *repository.PostRepository) *PostService {
	return &PostService{r}
}

// CreatePost - публикация нового поста
func (s *PostService) CreatePost(req dto.CreatePostRequest, authorID, authorTag string) (string, error) {
	ID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return "", fmt.Errorf("invalid author id")
	}

	postID, err := s.repository.CreatePost(req.Content, authorTag, ID)
	if err != nil {
		return "", fmt.Errorf("failed to create post: %w", err)
	}

	return postID, nil
}

// GetPost - получить пост по заданному id
func (s *PostService) GetPost(postID string) (model.Post, error) {
	ID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return model.Post{}, fmt.Errorf("invalid post id")
	}

	post, err := s.repository.GetPost(ID)
	if err != nil {
		return model.Post{}, fmt.Errorf("post not found: %w", err)
	}

	return post, nil
}

// GetAllPostsOfUser - получить все посты по id пользователя
func (s *PostService) GetAllPostsOfUser(authorID string) ([]model.Post, error) {
	ID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return nil, fmt.Errorf("invalid author id")
	}

	posts, err := s.repository.GetAllPostsOfUser(ID)
	if err != nil {
		return nil, fmt.Errorf("posts not found: %w", err)
	}

	return posts, nil
}

// SetLike - поставить лайк под заданным постом
func (s *PostService) SetLike(post, liker string) error {
	postID, err := primitive.ObjectIDFromHex(post)
	if err != nil {
		return fmt.Errorf("invalid post id")
	}
	likerID, err := primitive.ObjectIDFromHex(liker)
	if err != nil {
		return fmt.Errorf("invalid liker id")
	}

	if err = s.repository.SetLike(postID, likerID); err != nil {
		return fmt.Errorf("post cannot be liked")
	}

	return nil
}

// DeleteLike - убрать лайк под заданным постом
func (s *PostService) DeleteLike(post, liker string) error {
	postID, err := primitive.ObjectIDFromHex(post)
	if err != nil {
		return fmt.Errorf("invalid post id")
	}
	likerID, err := primitive.ObjectIDFromHex(liker)
	if err != nil {
		return fmt.Errorf("invalid liker id")
	}

	if err = s.repository.DeleteLike(postID, likerID); err != nil {
		return fmt.Errorf("post cannot be disliked")
	}

	return nil
}
