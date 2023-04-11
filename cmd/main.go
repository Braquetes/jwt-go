package main

import (
	"github.com/braquetes/jwt-go/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Use(handlers.Authorization)

	app.Get("/", handlers.Authorization, func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	app.Post("/generate", handlers.Generar)
	app.Post("/validate", handlers.Verificar)

	app.Listen(":3000")
}
