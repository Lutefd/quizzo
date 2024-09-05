package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Lutefd/quizzo/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "quizzes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", cfg.MongoUser, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort, cfg.MongoDBName)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(cfg.MongoDBName)

	collections, err := db.ListCollectionNames(context.Background(), bson.M{"name": collectionName})
	if err != nil {
		log.Fatalf("Failed to list collections: %v", err)
	}

	if len(collections) == 0 {
		err = db.CreateCollection(context.Background(), collectionName)
		if err != nil {
			log.Fatalf("Failed to create collection: %v", err)
		}
		fmt.Printf("Created '%s' collection in '%s' database\n", collectionName, cfg.MongoDBName)
	}

	fmt.Printf("Successfully connected to MongoDB\n")
	fmt.Printf("Database '%s' and collection '%s' are ready for use!\n", cfg.MongoDBName, collectionName)
}
