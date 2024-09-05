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
	os.Setenv("MONGO_DB_NAME", "testdb")
	os.Setenv("SERVER_PORT", "8089")

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
	if config.MongoDBName != "testdb" {
		t.Errorf("Expected MongoDBName to be 'testdb', got '%s'", config.MongoDBName)
	}
	if config.ServerPort != 8089 {
		t.Errorf("Expected ServerPort to be 8089, got '%d'", config.ServerPort)
	}
}

func TestLoadConfigMissingVars(t *testing.T) {
	os.Unsetenv("MONGO_HOST")
	os.Unsetenv("MONGO_USER")
	os.Unsetenv("MONGO_PASSWORD")
	os.Unsetenv("MONGO_PORT")
	os.Unsetenv("MONGO_DB_NAME")
	os.Unsetenv("SERVER_PORT")

	_, err := LoadConfig(true)
	if err == nil {
		t.Fatal("LoadConfig() should have failed with missing environment variables")
	}
}
