package handler

import (
	"project/services/post-service/internal/service"

	"github.com/gin-gonic/gin"
)

// PostHandler - структура,
// которая обрабатывает
// поступающий от клиента запрос
type PostHandler struct {
	s *service.PostService
}

// NewPostHandler - конструктор PostHandler
func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{s}
}

// CreatePost - публикация нового поста
func (h *PostHandler) CreatePost(c *gin.Context) {

}

// GetPost - получить пост по заданному id
func (h *PostHandler) GetPost(c *gin.Context) {

}

// GetAllPostsOfUser - получить все посты по id пользователя
func (h *PostHandler) GetAllPostsOfUser(c *gin.Context) {

}

// SetLike - поставить лайк под заданным постом
func (h *PostHandler) SetLike(c *gin.Context) {

}

// DeleteLike - убрать лайк под заданным постом
func (h *PostHandler) DeleteLike(c *gin.Context) {

}
