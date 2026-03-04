package handler

import "github.com/gin-gonic/gin"

// UserHandler - описание структуры UserHadler
type UserHandler struct{}

// NewUserHandler - создание нового объекта UserHandler
func NewUserHandler() *UserHandler{
	return &UserHandler{}
}

// GetUser - метод, возвращающий пользователя с заданным id
func (h *UserHandler)GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
		"message": "ok",
	})
}

// GetMe - метод, возвращающий свой профиль
func (h *UserHandler)GetMe(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

// RegisterUser - метод, осуществляющий регистрацию нового пользователя
func (h *UserHandler)RegisterUser(c *gin.Context) {
	c.JSON(201, gin.H{"message": "ok"})
}

// LoginUser - метод, возвращающий токен существующего пользователя
func (h *UserHandler)LoginUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

// GetFollowing - метод, возвращающий количество подписчиков
func (h *UserHandler)GetFollowing(c *gin.Context){
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
		"message": "ok",
	})
}

// SetFollow - метод, выполняющий подписку
func (h *UserHandler)SetFollow(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
		"message": "ok",
	})
}

// DeleteFollow - метод, выполняющий отписку
func (h *UserHandler)DeleteFollow(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
		"message": "ok",
	})
}

// UpdateAvatar - метод, обновляющий аватар пользователя
func (h *UserHandler)UpdateAvatar(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}