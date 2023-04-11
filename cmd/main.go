package main

import (
	"github.com/braquetes/jwt-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	app.Get("/generate", handlers.Generate)

	app.Listen(":3000")
}
