package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// UserClient - клиент для прокидывания запросов user-service
type UserClient struct {
	baseURL string
	client  *http.Client
}

// NewUserClient - конструктор UserClient
func NewUserClient(url string) *UserClient {
	return &UserClient{
		baseURL: url,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetFollowing - метод, обращающийся к user-service
// для получения списка подписок пользователя
func (c *UserClient) GetFollowing(token string) ([]string, error) {
	url := fmt.Sprintf("%s/users/me/following", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get following")
	}

	var result struct {
		Data struct{
			Followings []string `json:"followings"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data.Followings, nil
}
