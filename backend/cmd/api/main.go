package main

import "github.com/gofiber/fiber/v2"

func main() {
	srv := fiber.New()
	srv.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	srv.Listen(":8089")
}
