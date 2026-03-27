package dto

// CreatePostRequest - содержимое поста в запросе
// при создании поста
type CreatePostRequest struct {
	Content   string `json:"content" binding:"required,min=1,max=500"`
}
