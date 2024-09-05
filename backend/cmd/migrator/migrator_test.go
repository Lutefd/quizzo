package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/Lutefd/quizzo/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDBName = "quizzo_test"
const testCollectionName = "quizzes"

func setupTestDB(t *testing.T) (*mongo.Client, *mongo.Database, func()) {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", cfg.MongoUser, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort, testDBName)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db := client.Database(testDBName)

	collections, err := db.ListCollectionNames(context.Background(), bson.M{"name": testCollectionName})
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}

	if len(collections) == 0 {
		err = db.CreateCollection(context.Background(), testCollectionName)
		if err != nil {
			t.Fatalf("Failed to create test collection: %v", err)
		}
	}

	cleanup := func() {
		err := client.Database(testDBName).Drop(context.Background())
		if err != nil {
			t.Fatalf("Failed to drop test database: %v", err)
		}
		err = client.Disconnect(context.Background())
		if err != nil {
			t.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}

	return client, db, cleanup
}

func TestMongoDBConnection(t *testing.T) {
	_, db, cleanup := setupTestDB(t)
	defer cleanup()

	collections, err := db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}

	found := false
	for _, coll := range collections {
		if coll == testCollectionName {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("'%s' collection not found in test database", testCollectionName)
	}

	quizzesCollection := db.Collection(testCollectionName)
	_, err = quizzesCollection.InsertOne(context.Background(), bson.M{"name": "Test Quiz"})
	if err != nil {
		t.Fatalf("Failed to insert test document: %v", err)
	}

	var result bson.M
	err = quizzesCollection.FindOne(context.Background(), bson.M{"name": "Test Quiz"}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to retrieve test document: %v", err)
	}

	if result["name"] != "Test Quiz" {
		t.Errorf("Retrieved document does not match inserted document")
	}
}

func TestCreateCollection(t *testing.T) {
	_, db, cleanup := setupTestDB(t)
	defer cleanup()

	collections, err := db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}
	initialCount := len(collections)

	err = db.CreateCollection(context.Background(), testCollectionName)
	if err != nil {
		t.Errorf("Unexpected error when creating an existing collection: %v", err)
	}

	collections, err = db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}
	finalCount := len(collections)

	if finalCount != initialCount {
		t.Errorf("Expected %d collections, but got %d after attempting to create existing collection", initialCount, finalCount)
	}

	collections, err = db.ListCollectionNames(context.Background(), bson.M{"name": testCollectionName})
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}

	if len(collections) != 1 || collections[0] != testCollectionName {
		t.Errorf("Expected one collection named '%s', but got %v", testCollectionName, collections)
	}
}
