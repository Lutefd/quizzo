package repository

import (
	"testing"

	"github.com/Lutefd/quizzo/internal/config"
)

func TestNewMongoRepository(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	repo, err := NewMongoRepository(cfg)
	if err != nil {
		t.Fatalf("Failed to create MongoDB repository: %v", err)
	}
	if repo == nil {
		t.Fatal("Expected non-nil repository, got nil")
	}

	invalidCfg := &config.Config{
		MongoHost: "invalidhost",
		MongoPort: "invalidport",
	}
	_, err = NewMongoRepository(invalidCfg)
	if err == nil {
		t.Fatal("Expected error with invalid configuration, got nil")
	}
}
