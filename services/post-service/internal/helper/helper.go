package helper

import (
	"project/services/post-service/internal/dto"
	"project/services/post-service/internal/model"
	"time"
)

// ToSinglePostResponse - функция преобразования
// из внутренней модели поста в dto поста для
// передачи клиенту
func ToSinglePostResponse(postModel model.Post) dto.SinglePostResponse {
	return dto.SinglePostResponse{
		ID:         postModel.ID.Hex(),
		AuthorID:   postModel.AuthorID.Hex(),
		Content:    postModel.Content,
		LikesCount: postModel.LikesCount,
		CreatedAt:  postModel.CreatedAt.Format(time.RFC3339),
	}
}

// ToListPostResponse - функция преобразования списка постов
// в dto списка постов для передачи клиенту
func ToListPostResponse(list []model.Post) dto.ListPostResponse {
	data := make([]dto.SinglePostResponse, 0, len(list))
	for _, post := range list {
		data = append(data, ToSinglePostResponse(post))
	}

	return dto.ListPostResponse{Data: data}
}
