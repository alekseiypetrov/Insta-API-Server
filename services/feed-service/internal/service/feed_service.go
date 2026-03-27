package service

import "project/services/feed-service/internal/client"

// FeedService - сервис, выполняющий бизнес-логику feed-service
type FeedService struct {
	userClient *client.UserClient
	postClient *client.PostClient
}

// NewFeedService - конструктор FeedService
func NewFeedService(u *client.UserClient, p *client.PostClient) *FeedService {
	return &FeedService{
		userClient: u,
		postClient: p,
	}
}
