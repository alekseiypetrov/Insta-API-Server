package helper

import (
	"project/services/user-service/internal/dto"
	"project/services/user-service/internal/model"
)

// ToUserResponse - функция, преобразующая User из model
// в UserResponse для отправки клиенту
func ToUserResponse(user model.User) dto.UserResponse {
	response := dto.UserResponse{}
	response.ID = user.ID.Hex()
	response.Tag = user.Tag
	response.AvatarURL = user.AvatarURL
	response.FollowersCount = user.FollowersCount
	response.FollowingCount = user.FollowingCount
	return response
}
