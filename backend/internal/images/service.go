package images

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// ImageService interface for dependency injection
type ImageService interface {
	GetImages() ([]string, error)
	GetTags(imageName string) ([]string, error)
	RemoveImage(imageName string) error
	SearchImages(query string) ([]string, error)
	DeleteUnusedImages() error
	GetImageHistory(imageName string) ([]image.HistoryResponseItem, error)
}

type imageService struct {
	cli *client.Client
}

// NewImageService creates a new instance of ImageService using the Docker SDK
func NewImageService() (ImageService, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	// Optional: Set custom options based on config
	// For example, if your config has a specific host:
	// cli, err = client.NewClientWithOpts(client.WithHost(cfg.Registry.BaseURL))
	// if err != nil {
	//     return nil, err
	// }

	return &imageService{cli: cli}, nil
}

func (s *imageService) GetImages() ([]string, error) {
	images, err := s.cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		return nil, err
	}

	var imageNames []string
	for _, img := range images {
		for _, repoTag := range img.RepoTags {
			imageNames = append(imageNames, repoTag)
		}
	}
	return imageNames, nil
}

func (s *imageService) GetTags(imageName string) ([]string, error) {
	// With Docker SDK, we get tags when listing images.
	// This function needs adjustment depending on how you want to use it.
	// If you have the image ID, you can inspect the image to get the RepoTags.
	imageSummary, _, err := s.cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return nil, err
	}
	return imageSummary.RepoTags, nil

}

func (s *imageService) RemoveImage(imageName string) error {

	_, err := s.cli.ImageRemove(context.Background(), imageName, image.RemoveOptions{
		Force:         false, // Consider adding force option
		PruneChildren: true,
	})
	if err != nil {
		return err
	}
	log.Printf("Successfully removed image: %s", imageName)
	return nil
}

func (s *imageService) SearchImages(query string) ([]string, error) {
	images, err := s.GetImages()
	if err != nil {
		return nil, err
	}

	var filteredImages []string
	for _, image := range images {
		if containsCaseInsensitive(image, query) {
			filteredImages = append(filteredImages, image)
		}
	}
	return filteredImages, nil

}

func (s *imageService) DeleteUnusedImages() error {
	images, err := s.cli.ImageList(context.Background(), image.ListOptions{
		All: true, // Include dangling images
	})
	if err != nil {
		return err
	}

	for _, image := range images {
		if len(image.RepoTags) == 0 || (len(image.RepoTags) == 1 && image.RepoTags[0] == "<none>:<none>") { // Check for dangling images
			if err := s.RemoveImage(image.ID); err != nil { // Use image ID for removal
				log.Printf("Failed to delete image %s: %v", image.ID, err)
			} else {
				log.Printf("Deleted unused image: %s", image.ID)
			}
		}
	}
	return nil
}

func (s *imageService) GetImageHistory(imageName string) ([]image.HistoryResponseItem, error) {

	history, err := s.cli.ImageHistory(context.Background(), imageName)
	if err != nil {
		return nil, fmt.Errorf("failed to get image history: %w", err)
	}
	return history, nil
}

// Helper function for case-insensitive string match
func containsCaseInsensitive(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
