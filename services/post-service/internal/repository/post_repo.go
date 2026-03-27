package repository

import (
	"context"
	"errors"
	"fmt"
	"project/pkg/database"
	"project/services/post-service/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PostRepository - структура,
// которая выполняет запросы к БД
type PostRepository struct {
	collection *mongo.Collection
}

// NewPostRepository - конструктор репозитория
func NewPostRepository(connection *database.Mongo, collectionName string) *PostRepository {
	return &PostRepository{
		collection: connection.DB.Collection(collectionName),
	}
}

// CreatePost - публикация нового поста
func (r *PostRepository) CreatePost(content string, authorID primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newPost := model.Post{
		AuthorID:   authorID,
		Content:    content,
		LikesCount: 0,
		Likers:     []primitive.ObjectID{},
		CreatedAt:  time.Now(),
	}

	post, err := r.collection.InsertOne(ctx, newPost)
	if err != nil {
		return "", err
	}
	postID, ok := post.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to parse inserted post")
	}

	return postID.Hex(), nil
}

// GetPost - получить пост по заданному id
func (r *PostRepository) GetPost(postID primitive.ObjectID) (model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post model.Post
	filter := bson.M{"_id": postID}
	projection := bson.M{
		"likers": 0,
	}
	opts := options.FindOne().SetProjection(projection)
	err := r.collection.FindOne(ctx, filter, opts).Decode(&post)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Post{}, fmt.Errorf("post not found")
	}
	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

// GetAllPostsOfUser - получить все посты по id пользователя
func (r *PostRepository) GetAllPostsOfUser(userID primitive.ObjectID) ([]model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var postsList []model.Post
	filter := bson.M{"author_id": userID}
	projection := bson.M{
		"likers": 0,
	}
	opts := options.Find().SetProjection(projection)
	list, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if err = list.All(ctx, &postsList); err != nil {
		return nil, err
	}

	return postsList, nil
}

// SetLike - поставить лайк под заданным постом
func (r *PostRepository) SetLike(postID, likerID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"_id":    postID,
		"likers": bson.M{"$ne": likerID},
	}
	update := bson.D{
		{Key: "$addToSet", Value: bson.D{{Key: "likers", Value: likerID}}},
		{Key: "$inc", Value: bson.D{{Key: "likes_count", Value: 1}}},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("user has already liked this post")
	}

	return nil
}

// DeleteLike - убрать лайк под заданным постом
func (r *PostRepository) DeleteLike(postID, likerID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"_id":    postID,
		"likers": likerID,
	}
	update := bson.D{
		{Key: "$pull", Value: bson.D{{Key: "likers", Value: likerID}}},
		{Key: "$inc", Value: bson.D{{Key: "likes_count", Value: -1}}},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("user has already disliked this post")
	}

	return nil
}
