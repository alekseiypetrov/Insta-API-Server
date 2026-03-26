package service

import "project/services/post-service/internal/repository"

// PostService - структура, 
// выполняющая бизнес-логику микросервиса
type PostService struct {
	r *repository.PostRepository
}

// NewPostService - конструктор PostService
func NewPostService(r *repository.PostRepository) *PostService {
	return &PostService{r}
}
