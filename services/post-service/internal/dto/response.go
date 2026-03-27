package dto

// SinglePostResponse - ответ сервера из одного поста
type SinglePostResponse struct {
	ID         string `json:"id"`
	AuthorID   string `json:"author_id"`
	Content    string `json:"content"`
	LikesCount int    `bson:"likes_count"`
	CreatedAt  string `json:"created_at"`
}

// ListPostResponse - ответ сервера в виде списка постов
type ListPostResponse struct {
	Data []SinglePostResponse `json:"data"`
}
