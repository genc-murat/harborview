package images

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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
	BuildImage(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (string, error)
	PullImage(ctx context.Context, imageName, tag string) error
	PushImage(ctx context.Context, imageName string, options image.PushOptions) error
	SaveImage(ctx context.Context, imageNames []string, outputFile string) error
	PruneImages(ctx context.Context) (image.PruneReport, error)
	InspectImage(imageName string) (*types.ImageInspect, error)
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

func (s *imageService) BuildImage(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (string, error) {
	response, err := s.cli.ImageBuild(ctx, buildContext, options)
	if err != nil {
		return "", fmt.Errorf("failed to build image: %w", err)
	}
	defer response.Body.Close()

	// Parse response and return image ID
	var imageID string
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Successfully built") {
			parts := strings.Fields(line)
			imageID = parts[len(parts)-1]
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading build response: %w", err)
	}
	return imageID, nil
}

func (s *imageService) PullImage(ctx context.Context, imageName, tag string) error {
	imageRef := fmt.Sprintf("%s:%s", imageName, tag)
	response, err := s.cli.ImagePull(ctx, imageRef, image.PullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	}
	defer response.Close()

	// Print output for logging
	scanner := bufio.NewScanner(response)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	return nil
}

func (s *imageService) PushImage(ctx context.Context, imageName string, options image.PushOptions) error {
	response, err := s.cli.ImagePush(ctx, imageName, options)
	if err != nil {
		return fmt.Errorf("failed to push image: %w", err)
	}
	defer response.Close()

	// Print output for logging
	scanner := bufio.NewScanner(response)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	return nil
}

func (s *imageService) SaveImage(ctx context.Context, imageNames []string, outputFile string) error {
	response, err := s.cli.ImageSave(ctx, imageNames)
	if err != nil {
		return fmt.Errorf("failed to save images: %w", err)
	}
	defer response.Close()

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, response); err != nil {
		return fmt.Errorf("failed to write image tar: %w", err)
	}
	log.Printf("Images saved to %s", outputFile)
	return nil
}

func (s *imageService) PruneImages(ctx context.Context) (image.PruneReport, error) {
	report, err := s.cli.ImagesPrune(ctx, filters.Args{})
	if err != nil {
		return image.PruneReport{}, fmt.Errorf("failed to prune images: %w", err)
	}
	return report, nil
}

func (s *imageService) InspectImage(imageName string) (*types.ImageInspect, error) {
	inspect, _, err := s.cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect image: %w", err)
	}
	return &inspect, nil
}

// Helper function for case-insensitive string match
func containsCaseInsensitive(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
