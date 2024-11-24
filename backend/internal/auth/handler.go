package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/login", Login)
}

func Login(c *fiber.Ctx) error {
	// JWT token üretimi
	return c.JSON(fiber.Map{"token": "JWT_TOKEN_HERE"})
}
