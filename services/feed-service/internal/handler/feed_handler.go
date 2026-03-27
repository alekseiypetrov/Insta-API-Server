package handler

import (
	"project/services/feed-service/internal/service"

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

}
