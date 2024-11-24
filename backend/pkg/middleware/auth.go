package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	// Auth kontrol
	return c.Next()
}
