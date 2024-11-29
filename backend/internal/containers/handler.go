package containers

import (
	"io"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all routes for containers
func RegisterRoutes(app *fiber.App) {
	// Initialize the service
	service, _ := NewContainerService()

	containersGroup := app.Group("/containers")

	// List all containers
	containersGroup.Get("/", func(c *fiber.Ctx) error {
		return ListContainers(c, service)
	})

	// Inspect a specific container
	containersGroup.Get("/:id", func(c *fiber.Ctx) error {
		return InspectContainer(c, service)
	})

	// Start a container
	containersGroup.Post("/:id/start", func(c *fiber.Ctx) error {
		return StartContainer(c, service)
	})

	// Stop a container
	containersGroup.Post("/:id/stop", func(c *fiber.Ctx) error {
		return StopContainer(c, service)
	})

	// Restart a container
	containersGroup.Post("/:id/restart", func(c *fiber.Ctx) error {
		return RestartContainer(c, service)
	})

	// Remove a container
	containersGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return RemoveContainer(c, service)
	})

	// Get logs for a container
	containersGroup.Get("/:id/logs", func(c *fiber.Ctx) error {
		return GetContainerLogs(c, service)
	})

	// Execute a command in a container
	containersGroup.Post("/:id/exec", func(c *fiber.Ctx) error {
		return ExecInContainer(c, service)
	})

	// Get stats for a container
	containersGroup.Get("/:id/stats", func(c *fiber.Ctx) error {
		return GetContainerStats(c, service)
	})

	// Prune unused containers
	containersGroup.Post("/prune", func(c *fiber.Ctx) error {
		return PruneContainers(c, service)
	})
}

// ListContainers handler for fetching all containers
func ListContainers(c *fiber.Ctx, service ContainerService) error {
	all := c.QueryBool("all", false) // Default: show only running containers
	containers, err := service.ListContainers(all)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(containers)
}

// InspectContainer handler for inspecting a specific container
func InspectContainer(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	container, err := service.InspectContainer(containerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(container)
}

// StartContainer handler for starting a container
func StartContainer(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	err := service.StartContainer(containerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container started successfully"})
}

// StopContainer handler for stopping a container
func StopContainer(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	timeout := c.QueryInt("timeout", 10) // Default timeout: 10 seconds
	err := service.StopContainer(containerID, &timeout)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container stopped successfully"})
}

// RestartContainer handler for restarting a container
func RestartContainer(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	timeout := c.QueryInt("timeout", 10) // Default timeout: 10 seconds
	err := service.RestartContainer(containerID, &timeout)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container restarted successfully"})
}

// RemoveContainer handler for deleting a container
func RemoveContainer(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	force := c.QueryBool("force", false) // Default: do not force
	err := service.RemoveContainer(containerID, force)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Container removed successfully"})
}

// GetContainerLogs handler for fetching logs for a container
func GetContainerLogs(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	follow := c.QueryBool("follow", false)
	logs, err := service.ContainerLogs(containerID, follow)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer logs.Close()

	data, _ := io.ReadAll(logs)
	return c.JSON(fiber.Map{"logs": string(data)})
}

// ExecInContainer handler for executing a command in a container
func ExecInContainer(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	var request struct {
		Cmd []string `json:"cmd"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	output, err := service.ExecInContainer(containerID, request.Cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"output": output})
}

// GetContainerStats handler for fetching stats for a container
func GetContainerStats(c *fiber.Ctx, service ContainerService) error {
	containerID := c.Params("id")
	stream := c.QueryBool("stream", false)
	stats, err := service.Stats(containerID, stream)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer stats.Close()

	data, _ := io.ReadAll(stats)
	return c.JSON(fiber.Map{"stats": string(data)})
}

// PruneContainers handler for pruning unused containers
func PruneContainers(c *fiber.Ctx, service ContainerService) error {
	report, err := service.PruneContainers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(report)
}
