package repository

import (
	"user-service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository - структура, представляющая собой репозиторий,
// который является посредником между service и БД
type UserRepository struct {
	db *mongo.Client
}

// NewUserRepository - конструктор репозитория
func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// QueryUser - функция, выполняющая поиск пользователя по тегу
func (r *UserRepository) QueryUser(tag string) (model.LoginUser, bool) {
	return model.LoginUser{}, true
}

// CreateUser - функция, выполняющая создание нового пользователя
func (r *UserRepository) CreateUser(tag, password, answer string) (string, error) {
	return "", nil
}

// QueryProfile - функция, выполняющая поиск пользователя по id
func (r *UserRepository) QueryProfile(id primitive.ObjectID) (model.User, error) {
	return model.User{}, nil
}

// InsertFollow - функция, выполняющая оформление подписки
func (r *UserRepository) InsertFollow() {}

// DeleteFollow - функция, выполняющая удаление подписки
func (r *UserRepository) DeleteFollow() {}

// QueryFollowers - функция, выполняющая поиск подписчиков
func (r *UserRepository) QueryFollowers() {}

// UpdateAvatar - функция, выполняющая обновление аватара пользователя
func (r *UserRepository) UpdateAvatar() {}
