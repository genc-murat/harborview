package images

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/genc-murat/harborview/pkg/config"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all routes for images
func RegisterRoutes(app *fiber.App, cfg config.Config) {
	// Initialize the service with the provided config
	service, _ := NewImageService()

	imagesGroup := app.Group("/images")

	// Existing routes
	imagesGroup.Get("/", func(c *fiber.Ctx) error {
		return GetImages(c, service)
	})

	imagesGroup.Get("/:name/tags", func(c *fiber.Ctx) error {
		return GetTags(c, service)
	})

	imagesGroup.Delete("/:name", func(c *fiber.Ctx) error {
		return RemoveImage(c, service)
	})

	imagesGroup.Get("/search", func(c *fiber.Ctx) error {
		return SearchImages(c, service)
	})

	imagesGroup.Get("/:name/history", func(c *fiber.Ctx) error {
		return GetImageHistory(c, service)
	})

	imagesGroup.Delete("/unused", func(c *fiber.Ctx) error {
		return DeleteUnusedImages(c, service)
	})

	// New routes
	imagesGroup.Post("/build", func(c *fiber.Ctx) error {
		return BuildImage(c, service)
	})

	imagesGroup.Post("/pull", func(c *fiber.Ctx) error {
		return PullImage(c, service)
	})

	imagesGroup.Post("/push/:name", func(c *fiber.Ctx) error {
		return PushImage(c, service)
	})

	imagesGroup.Post("/save", func(c *fiber.Ctx) error {
		return SaveImage(c, service)
	})

	imagesGroup.Post("/prune", func(c *fiber.Ctx) error {
		return PruneImages(c, service)
	})

	imagesGroup.Get("/:name/inspect", func(c *fiber.Ctx) error {
		return InspectImage(c, service)
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

func BuildImage(c *fiber.Ctx, service ImageService) error {
	// Parse build context from the request body
	buildContext := c.Body() // Assuming the build context is sent in the request body

	// Parse tags query parameter into a slice of strings
	tagsQuery := c.Query("tags", "")
	var tags []string
	if tagsQuery != "" {
		tags = strings.Split(tagsQuery, ",")
	}

	options := types.ImageBuildOptions{
		Tags: tags, // Use the parsed slice of tags
	}

	// Call the BuildImage service method
	imageID, err := service.BuildImage(c.Context(), strings.NewReader(string(buildContext)), options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"imageID": imageID,
	})
}

func PullImage(c *fiber.Ctx, service ImageService) error {
	imageName := c.Query("name")
	tag := c.Query("tag", "latest")
	err := service.PullImage(c.Context(), imageName, tag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Image pulled successfully"})
}

func PushImage(c *fiber.Ctx, service ImageService) error {
	imageName := c.Params("name")
	options := image.PushOptions{}
	err := service.PushImage(c.Context(), imageName, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Image pushed successfully"})
}

func SaveImage(c *fiber.Ctx, service ImageService) error {
	var request struct {
		Images []string `json:"images"`
		Output string   `json:"output"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	err := service.SaveImage(c.Context(), request.Images, request.Output)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Images saved successfully"})
}

func PruneImages(c *fiber.Ctx, service ImageService) error {
	report, err := service.PruneImages(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(report)
}

func InspectImage(c *fiber.Ctx, service ImageService) error {
	imageName := c.Params("name")
	inspect, err := service.InspectImage(imageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(inspect)
}
