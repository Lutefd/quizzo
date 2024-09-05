package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Lutefd/quizzo/internal/collection"
	"github.com/Lutefd/quizzo/internal/commons"
	"github.com/Lutefd/quizzo/internal/config"
	"github.com/Lutefd/quizzo/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config *config.Config
	app    *fiber.App
	repo   collection.QuizzesCollection
}

func NewServer(cfg *config.Config, initRepo ...bool) (*Server, error) {
	repo, err := repository.NewMongoRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize repository: %w", err)
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Duration(commons.ServerIdleTimeout),
		ReadTimeout:  time.Duration(commons.ServerReadTimeout),
		WriteTimeout: time.Duration(commons.ServerWriteTimeout),
	})

	server := &Server{
		config: cfg,
		app:    app,
		repo:   repo,
	}

	server.setupRoutes()

	return server, nil
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		fmt.Printf("Server started on port %d\n", s.config.ServerPort)
		if err := s.app.Listen(fmt.Sprintf(":%d", s.config.ServerPort)); err != nil {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	<-ctx.Done()
	return s.Shutdown()
}

func (s *Server) Shutdown() error {
	fmt.Println("Server is shutting down...")

	if err := s.app.Shutdown(); err != nil {
		fmt.Printf("HTTP server shutdown error: %v\n", err)
		return err
	}

	if err := s.repo.Close(); err != nil {
		fmt.Printf("Repository close error: %v\n", err)
		return err
	}
	fmt.Println("Server shutdown complete")
	return nil
}
