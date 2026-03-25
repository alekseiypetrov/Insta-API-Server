package repository

import (
	"context"
	"fmt"
	"project/services/user-service/internal/model"
	"time"

	"project/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName string = "users"
)

// UserRepository - репозиторий, выполняющий запросы к БД
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository - конструктор репозитория
func NewUserRepository(db *database.Mongo) *UserRepository {
	return &UserRepository{
		collection: db.DB.Collection(collectionName),
	}
}

// CreateIndexes - создание уникального индекса на поля
func (r *UserRepository) CreateIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "tag", Value: 1}},
		Options: options.Index().
			SetUnique(true),
	}

	_, err := r.collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

// CreateUser - функция, выполняющая создание нового пользователя
func (r *UserRepository) CreateUser(tag, password, answer string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newUser := model.User{
		Tag:            tag,
		Password:       password,
		SecretAnswer:   answer,
		AvatarURL:      "",
		FollowingCount: 0,
		FollowersCount: 0,
		Followers:      []primitive.ObjectID{},
		Following:      []primitive.ObjectID{},
	}

	user, err := r.collection.InsertOne(ctx, newUser)
	if err != nil {
		return "", err
	}

	id, ok := user.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to parse inserted id")
	}

	return id.Hex(), nil
}

// QueryUser - функция, выполняющая поиск пользователя по тегу
func (r *UserRepository) QueryUser(tag string) (model.LoginUser, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result model.LoginUser
	filter := bson.M{"tag": tag}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return model.LoginUser{}, false
	}
	return result, true
}

// QueryProfile - функция, выполняющая поиск пользователя по id
func (r *UserRepository) QueryProfile(id primitive.ObjectID) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result model.User
	filter := bson.M{"_id": id}
	projection := bson.M{
		"password_hash": 0,
		"secret_answer": 0,
	}
	opts := options.FindOne().SetProjection(projection)

	err := r.collection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return model.User{}, err
	}
	return result, nil
}

// InsertFollow - функция, выполняющая оформление подписки
func (r *UserRepository) InsertFollow(followerID, targetID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		filter primitive.M
		update primitive.D
	)

	filter = bson.M{
		"_id":       followerID,
		"following": bson.M{"$ne": targetID},
	}
	update = bson.D{
		{Key: "$addToSet", Value: bson.D{{Key: "following", Value: targetID}}},
		{Key: "$inc", Value: bson.D{{Key: "following_count", Value: 1}}},
	}
	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("already following or user not found")
	}

	filter = bson.M{
		"_id":       targetID,
		"followers": bson.M{"$ne": followerID},
	}
	update = bson.D{
		{Key: "$addToSet", Value: bson.D{{Key: "followers", Value: followerID}}},
		{Key: "$inc", Value: bson.D{{Key: "followers_count", Value: 1}}},
	}
	res, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("already followed or user not found")
	}

	return nil
}

// DeleteFollow - функция, выполняющая удаление подписки
func (r *UserRepository) DeleteFollow(followerID, targetID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		filter primitive.M
		update primitive.D
	)

	filter = bson.M{"_id": followerID, "following": targetID}
	update = bson.D{
		{Key: "$pull", Value: bson.D{{Key: "following", Value: targetID}}},
		{Key: "$inc", Value: bson.D{{Key: "following_count", Value: -1}}},
	}
	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("already unfollowing or user not found")
	}

	filter = bson.M{"_id": targetID, "followers": followerID}
	update = bson.D{
		{Key: "$pull", Value: bson.D{{Key: "followers", Value: followerID}}},
		{Key: "$inc", Value: bson.D{{Key: "followers_count", Value: -1}}},
	}
	res, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("already unfollowed or user not found")
	}

	return nil
}

// TODO: - Will be done later

// UpdateAvatar - функция, выполняющая обновление аватара пользователя
// func (r *UserRepository) UpdateAvatar() {}
