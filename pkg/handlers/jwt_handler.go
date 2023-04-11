package handlers

import (
	"strings"

	"github.com/braquetes/jwt-go/pkg/domain"
	"github.com/braquetes/jwt-go/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Generar(c *fiber.Ctx) error {
	jwt := new(domain.JWT_Params)
	if err := c.BodyParser(jwt); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Invalid body parser",
		})
	}
	if jwt.Username == "" || jwt.Email == "" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Username and Email is required",
		})
	}
	res, err := middlewares.GenerateToken(jwt)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": res,
	})
}

func Verificar(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]
	token, err := middlewares.ValidateToken(tokenString)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if token == nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token invalido",
		})
	}
	return c.Status(200).JSON(token)
}

func Authorization(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]
	token, err := middlewares.ValidateToken(tokenString)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if token == nil {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Token invalido",
		})
	}

	return c.Next()
}
