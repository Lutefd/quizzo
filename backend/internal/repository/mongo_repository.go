package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Lutefd/quizzo/internal/collection"
	"github.com/Lutefd/quizzo/internal/config"
	"github.com/Lutefd/quizzo/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(cfg *config.Config) (collection.QuizzesCollection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", cfg.MongoUser, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort, cfg.MongoDBName)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	database := client.Database(cfg.MongoDBName)
	collection := database.Collection("quizzes")

	return &MongoRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *MongoRepository) InsertQuiz(quiz model.Quiz) error {
	// Implementation to be added
	return nil
}

func (r *MongoRepository) GetQuizzes() ([]model.Quiz, error) {
	// Implementation to be added
	return nil, nil
}

func (r *MongoRepository) GetQuizById(id primitive.ObjectID) (*model.Quiz, error) {
	// Implementation to be added
	return nil, nil
}

func (r *MongoRepository) UpdateQuiz(quiz model.Quiz) error {
	// Implementation to be added
	return nil
}

func (r *MongoRepository) DeleteQuiz(id primitive.ObjectID) error {
	// Implementation to be added
	return nil
}

func (r *MongoRepository) Close() error {
	return r.client.Disconnect(context.Background())
}
