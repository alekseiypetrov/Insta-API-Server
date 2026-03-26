package dto

// LoginRequest - модель запроса
// для входа в существующий аккаунт
type LoginRequest struct {
	Tag      string `json:"tag" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest - модель запроса
// для регистрации нового пользователя
type RegisterRequest struct {
	Tag          string `json:"tag" binding:"required,min=2"`
	Password     string `json:"password" binding:"required,min=8,max=40"`
	SecretAnswer string `json:"secret_answer" binding:"required,min=4,max=40"`
}
