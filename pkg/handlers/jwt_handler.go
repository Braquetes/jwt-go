package handlers

import (
	"github.com/braquetes/jwt-go/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Generar(c *fiber.Ctx) error {
	res, err := middlewares.GenerateToken("userId")
	if err != nil {
		return err
	}
	return nil
}
