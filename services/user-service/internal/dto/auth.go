package dto

// LoginRequest - модель запроса
// для входа в существующий аккаунт
type LoginRequest struct {
	Tag      string `json:"tag" binding:"required,tag"`
	Password string `json:"password" binding:"required,password"`
}

// RegisterRequest - модель запроса
// для регистрации нового пользователя
type RegisterRequest struct {
	Tag          string `json:"tag" binding:"required,tag"`
	Password     string `json:"password" binding:"required,password"`
	SecretAnswer string `json:"secret_answer" binding:"required,secret_answer"`
}
