package handler

import (
	"user-service/internal/dto"
	"user-service/internal/helper"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler - описание структуры UserHadler
type UserHandler struct {
	authService *service.AuthService
	userService *service.UserService
}

// NewUserHandler - создание нового объекта UserHandler
func NewUserHandler(a *service.AuthService, u *service.UserService) *UserHandler {
	return &UserHandler{
		authService: a,
		userService: u,
	}
}

// GetUser - метод, возвращающий пользователя с заданным id
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	response := helper.ToUserResponse(user)
	c.JSON(200, gin.H{"data": response})
}

// GetMe - метод, возвращающий свой профиль
func (h *UserHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	id, ok := userID.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	response := helper.ToUserResponse(user)
	c.JSON(200, gin.H{"data": response})
}

// RegisterUser - метод, осуществляющий регистрацию нового пользователя
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.SignUp(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"token": token})
}

// LoginUser - метод, возвращающий токен существующего пользователя
func (h *UserHandler) LoginUser(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.SignIn(req)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

// SetFollow - метод, выполняющий подписку
func (h *UserHandler) SetFollow(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	followerID, ok := id.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	targetID := c.Param("id")
	if _, err := h.userService.GetByID(targetID); err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	if err := h.userService.FollowUser(followerID, targetID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// DeleteFollow - метод, выполняющий отписку
func (h *UserHandler) DeleteFollow(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	followerID, ok := id.(string)
	if !ok {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	targetID := c.Param("id")
	if _, err := h.userService.GetByID(targetID); err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	if err := h.userService.UnfollowUser(followerID, targetID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// TODO: - Will be done later

// UpdateAvatar - метод, обновляющий аватар пользователя
// func (h *UserHandler) UpdateAvatar(c *gin.Context) {
// 	c.JSON(200, gin.H{"message": "ok"})
// }

// GetFollowing - метод, возвращающий количество подписчиков
// func (h *UserHandler) GetFollowing(c *gin.Context) {
// 	id := c.Param("id")
// 	c.JSON(200, gin.H{
// 		"id":      id,
// 		"message": "ok",
// 	})
// }
