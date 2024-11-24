package images

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/genc-murat/harborview/pkg/config"
)

// ImageService interface for dependency injection
// ImageService interface for dependency injection
type ImageService interface {
	GetImages() ([]string, error)
	GetTags(imageName string) ([]string, error)
	RemoveImage(imageName string) error
	SearchImages(query string) ([]string, error)
	DeleteUnusedImages() error
	GetImageHistory(imageName string) (map[string]interface{}, error)
}

type imageService struct {
	baseURL  string
	username string
	password string
}

// NewImageService creates a new instance of ImageService
func NewImageService(cfg config.Config) ImageService {
	return &imageService{
		baseURL:  cfg.Registry.BaseURL,
		username: cfg.Registry.Username,
		password: cfg.Registry.Password,
	}
}

func (s *imageService) RemoveImage(imageName string) error {
	url := s.baseURL + "/v2/" + imageName + "/manifests/latest" // Example endpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Add Basic Auth if credentials are provided
	if s.username != "" && s.password != "" {
		req.SetBasicAuth(s.username, s.password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		log.Printf("Failed to remove image %s. Status code: %d", imageName, resp.StatusCode)
		return errors.New("failed to remove image")
	}

	log.Printf("Successfully removed image: %s", imageName)
	return nil
}

func (s *imageService) SearchImages(query string) ([]string, error) {
	images, err := s.GetImages()
	if err != nil {
		return nil, err
	}

	// Filter images based on query
	var filteredImages []string
	for _, image := range images {
		if containsCaseInsensitive(image, query) {
			filteredImages = append(filteredImages, image)
		}
	}
	return filteredImages, nil
}

// Helper function for case-insensitive string match
func containsCaseInsensitive(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func (s *imageService) DeleteUnusedImages() error {
	images, err := s.GetImages()
	if err != nil {
		return err
	}

	for _, image := range images {
		tags, err := s.GetTags(image)
		if err != nil {
			log.Printf("Failed to fetch tags for image %s: %v", image, err)
			continue
		}

		// If the image has no tags, delete it
		if len(tags) == 0 {
			if err := s.RemoveImage(image); err != nil {
				log.Printf("Failed to delete image %s: %v", image, err)
			} else {
				log.Printf("Deleted unused image: %s", image)
			}
		}
	}
	return nil
}

func (s *imageService) GetImageHistory(imageName string) (map[string]interface{}, error) {
	url := s.baseURL + "/v2/" + imageName + "/manifests/latest"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add Basic Auth if credentials are provided
	if s.username != "" && s.password != "" {
		req.SetBasicAuth(s.username, s.password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get history for image %s. Status code: %d", imageName, resp.StatusCode)
		return nil, errors.New("failed to get image history")
	}

	// Decode JSON response
	var history map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		return nil, err
	}

	return history, nil
}

// GetImages fetches the list of images from the Docker Registry
func (s *imageService) GetImages() ([]string, error) {
	url := s.baseURL + "/v2/_catalog"

	// HTTP request to Docker Registry
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add Basic Auth if credentials are provided
	if s.username != "" && s.password != "" {
		req.SetBasicAuth(s.username, s.password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		log.Printf("Docker Registry returned status code: %d", resp.StatusCode)
		return nil, errors.New("failed to fetch images")
	}

	// Decode JSON response
	var response struct {
		Repositories []string `json:"repositories"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Repositories, nil
}

// GetTags fetches the list of tags for a given image from the Docker Registry
func (s *imageService) GetTags(imageName string) ([]string, error) {
	url := s.baseURL + "/v2/" + imageName + "/tags/list"

	// HTTP request to Docker Registry
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add Basic Auth if credentials are provided
	if s.username != "" && s.password != "" {
		req.SetBasicAuth(s.username, s.password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		log.Printf("Docker Registry returned status code: %d", resp.StatusCode)
		return nil, errors.New("failed to fetch tags")
	}

	// Decode JSON response
	var response struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Tags, nil
}
