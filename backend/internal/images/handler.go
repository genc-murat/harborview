package images

import (
	"github.com/genc-murat/harborview/pkg/config"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all routes for images
func RegisterRoutes(app *fiber.App, cfg config.Config) {
	// Initialize the service with the provided config
	service := NewImageService(cfg)

	imagesGroup := app.Group("/images")

	// List all images
	imagesGroup.Get("/", func(c *fiber.Ctx) error {
		return GetImages(c, service)
	})

	// List tags for a specific image
	imagesGroup.Get("/:name/tags", func(c *fiber.Ctx) error {
		return GetTags(c, service)
	})

	// Remove a specific image
	imagesGroup.Delete("/:name", func(c *fiber.Ctx) error {
		return RemoveImage(c, service)
	})

	// Search for images
	imagesGroup.Get("/search", func(c *fiber.Ctx) error {
		return SearchImages(c, service)
	})

	// Get history of a specific image
	imagesGroup.Get("/:name/history", func(c *fiber.Ctx) error {
		return GetImageHistory(c, service)
	})

	// Delete unused images
	imagesGroup.Delete("/unused", func(c *fiber.Ctx) error {
		return DeleteUnusedImages(c, service)
	})
}

// GetImages handler for fetching all images
func GetImages(c *fiber.Ctx, service ImageService) error {
	images, err := service.GetImages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(images)
}

// GetTags handler for fetching tags for an image
func GetTags(c *fiber.Ctx, service ImageService) error {
	imageName := c.Params("name")
	tags, err := service.GetTags(imageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(tags)
}

// RemoveImage handler for deleting an image
func RemoveImage(c *fiber.Ctx, service ImageService) error {
	imageName := c.Params("name")
	err := service.RemoveImage(imageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Image removed successfully"})
}

// SearchImages handler for searching images by a query
func SearchImages(c *fiber.Ctx, service ImageService) error {
	query := c.Query("q", "")
	images, err := service.SearchImages(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(images)
}

// GetImageHistory handler for fetching the history of an image
func GetImageHistory(c *fiber.Ctx, service ImageService) error {
	imageName := c.Params("name")
	history, err := service.GetImageHistory(imageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(history)
}

// DeleteUnusedImages handler for deleting unused images
func DeleteUnusedImages(c *fiber.Ctx, service ImageService) error {
	err := service.DeleteUnusedImages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"message": "Unused images deleted successfully"})
}
