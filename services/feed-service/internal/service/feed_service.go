package service

import (
	"project/services/feed-service/internal/client"
	"project/services/feed-service/internal/dto"
	"sort"
)

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

// GetFeed - метод, возвращающий ленту по подпискам
func (s *FeedService) GetFeed(token string) ([]dto.SinglePostResponse, error) {
	following, err := s.userClient.GetFollowing(token)
	if err != nil {
		return nil, err
	}

	var feed = []dto.SinglePostResponse{}
	for _, id := range following {
		posts, err := s.postClient.GetPostsByUser(id)
		if err != nil {
			continue
		}

		feed = append(feed, posts...)
	}

	sort.Slice(feed, func(i, j int) bool {
		return feed[i].CreatedAt < feed[j].CreatedAt
	})

	return feed, nil
}
