package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Lutefd/quizzo/internal/config"
	"github.com/Lutefd/quizzo/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		fmt.Printf("Failed to create server: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := srv.Start(ctx); err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}
