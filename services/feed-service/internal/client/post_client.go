package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/services/feed-service/internal/dto"
	"time"
)

// PostClient - клиент для прокидывания запросов post-service
type PostClient struct {
	baseURL string
	client  *http.Client
}

// NewPostClient - конструктор PostClient
func NewPostClient(url string) *PostClient {
	return &PostClient{
		baseURL: url,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetPostsByUser - метод, обращающийся к post-service
// для получения постов пользователя
func (c *PostClient) GetPostsByUser(userID string) ([]dto.SinglePostResponse, error) {
	url := fmt.Sprintf("%s/users/%s/posts", c.baseURL, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get posts")
	}

	var result struct {
		Data []dto.SinglePostResponse `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
