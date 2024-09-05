package server

import "github.com/gofiber/fiber/v2"

func (s *Server) setupRoutes() {
	api := s.app.Group("/api")

	api.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
