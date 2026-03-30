package handler

import (
	"project/services/post-service/internal/dto"
	"project/services/post-service/internal/helper"
	"project/services/post-service/internal/service"

	"github.com/gin-gonic/gin"
)

// PostHandler - структура,
// которая обрабатывает
// поступающий от клиента запрос
type PostHandler struct {
	service *service.PostService
}

// NewPostHandler - конструктор PostHandler
func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{s}
}

// CreatePost - публикация нового поста
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	authorID, ok := id.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	tag, exists := c.Get("tag")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	authorTag, ok := tag.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	postID, err := h.service.CreatePost(req, authorID, authorTag)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	c.JSON(201, gin.H{
		"post_id": postID,
		"status":  "created",
	})
}

// GetPost - получить пост по заданному id
func (h *PostHandler) GetPost(c *gin.Context) {
	postID := c.Param("id")

	post, err := h.service.GetPost(postID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := helper.ToSinglePostResponse(post)
	c.JSON(200, gin.H{"data": response})
}

// GetAllPostsOfUser - получить все посты по id пользователя
func (h *PostHandler) GetAllPostsOfUser(c *gin.Context) {
	userID := c.Param("id")

	posts, err := h.service.GetAllPostsOfUser(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := helper.ToListPostResponse(posts)
	c.JSON(200, response)
}

// SetLike - поставить лайк под заданным постом
func (h *PostHandler) SetLike(c *gin.Context) {
	postID := c.Param("id")

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	likerID, ok := id.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	if err := h.service.SetLike(postID, likerID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

// DeleteLike - убрать лайк под заданным постом
func (h *PostHandler) DeleteLike(c *gin.Context) {
	postID := c.Param("id")

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	likerID, ok := id.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	if err := h.service.DeleteLike(postID, likerID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
