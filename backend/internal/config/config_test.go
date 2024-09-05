package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("MONGO_HOST", "testhost:27017")
	os.Setenv("MONGO_USERNAME", "testuser")
	os.Setenv("MONGO_PASSWORD", "testpassword")
	os.Setenv("MONGO_PORT", "27017")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if config.MongoHost != "testhost:27017" {
		t.Errorf("Expected MongoHost to be 'testhost:27017', got '%s'", config.MongoHost)
	}
	if config.MongoUser != "testuser" {
		t.Errorf("Expected MongoUser to be 'testuser', got '%s'", config.MongoUser)
	}
	if config.MongoPassword != "testpassword" {
		t.Errorf("Expected MongoPassword to be 'testpassword', got '%s'", config.MongoPassword)
	}
	if config.MongoPort != "27017" {
		t.Errorf("Expected MongoPort to be '27017', got '%s'", config.MongoPort)
	}
}

func TestLoadConfigMissingVars(t *testing.T) {
	os.Unsetenv("MONGO_HOST")
	os.Unsetenv("MONGO_USER")
	os.Unsetenv("MONGO_PASSWORD")
	os.Unsetenv("MONGO_PORT")

	_, err := LoadConfig(true)
	if err == nil {
		t.Fatal("LoadConfig() should have failed with missing environment variables")
	}
}
