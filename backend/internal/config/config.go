package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoHost     string
	MongoUser     string
	MongoPassword string
	MongoPort     string
}

func LoadConfig(skipEnvFile ...bool) (*Config, error) {
	if len(skipEnvFile) == 0 || !skipEnvFile[0] {
		err := godotenv.Load("../../../.env")
		if err != nil {
			fmt.Println("Error loading .env file, falling back to system environment variables")
		}
	}

	config := &Config{
		MongoHost:     os.Getenv("MONGO_HOST"),
		MongoUser:     os.Getenv("MONGO_USERNAME"),
		MongoPassword: os.Getenv("MONGO_PASSWORD"),
		MongoPort:     os.Getenv("MONGO_PORT"),
	}

	if config.MongoHost == "" || config.MongoUser == "" || config.MongoPassword == "" {
		return nil, fmt.Errorf("MONGO_HOST, MONGO_USERNAME, MONGO_PASSWORD AND MONGO_PORT must be set")
	}

	return config, nil
}
