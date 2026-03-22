package dto

// UserResponse - DTO для ответа клиенту
// с публичной информацией пользователя
type UserResponse struct {
	ID             string `json:"id"`
	Tag            string `json:"tag"`
	AvatarURL      string `json:"avatar_url"`
	FollowingCount int    `json:"following_count"`
	FollowersCount int    `json:"followers_count"`
}