package main

import (
	"log"

	"github.com/genc-murat/harborview/internal/auth"
	"github.com/genc-murat/harborview/internal/images"
	"github.com/genc-murat/harborview/pkg/config"
	"github.com/genc-murat/harborview/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                   // Tüm origin'lere izin ver (geliştirme için)
		AllowMethods: "GET,POST,PUT,DELETE", // İzin verilen HTTP metodları
	}))
	app.Use(middleware.AuthMiddleware)

	// Routes
	auth.RegisterRoutes(app)
	images.RegisterRoutes(app, cfg)

	// Start server
	log.Printf("Server starting on %s...", cfg.Server.Port)
	if err := app.Listen(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
