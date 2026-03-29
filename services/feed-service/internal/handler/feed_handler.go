package handler

import (
	"project/services/feed-service/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

// FeedHandler - обработчик запросов для feed-service
type FeedHandler struct {
	service *service.FeedService
}

// NewFeedHandler - конструктор FeedHandler
func NewFeedHandler(s *service.FeedService) *FeedHandler {
	return &FeedHandler{
		service: s,
	}
}

// GetFeed - получить ленту
func (h *FeedHandler) GetFeed(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	feed, err := h.service.GetFeed(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": feed})
}
