package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBName   string
	MongoHost     string
	MongoUser     string
	MongoPassword string
	MongoPort     string
	ServerPort    int
}

func LoadConfig(skipEnvFile ...bool) (*Config, error) {
	if len(skipEnvFile) == 0 || !skipEnvFile[0] {
		err := godotenv.Load("../../../.env")
		if err != nil {
			fmt.Println("Error loading .env file, falling back to system environment variables")
		}
	}
	srvPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("SERVER_PORT environment variable is not a valid integer")
	}
	config := &Config{
		MongoHost:     os.Getenv("MONGO_HOST"),
		MongoUser:     os.Getenv("MONGO_USERNAME"),
		MongoPassword: os.Getenv("MONGO_PASSWORD"),
		MongoPort:     os.Getenv("MONGO_PORT"),
		MongoDBName:   os.Getenv("MONGO_DB_NAME"),
		ServerPort:    srvPort,
	}

	if config.MongoHost == "" || config.MongoUser == "" || config.MongoPassword == "" || config.MongoPort == "" || config.MongoDBName == "" || config.ServerPort == 0 {
		return nil, fmt.Errorf("MONGO_HOST, MONGO_USERNAME, MONGO_PASSWORD, MONGO_PORT, MONGO_DB_NAME environment variables are required")
	}

	return config, nil
}
