package containers

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// ContainerService interface for dependency injection
type ContainerService interface {
	ListContainers(all bool) ([]types.Container, error)
	InspectContainer(containerID string) (*types.ContainerJSON, error)
	StartContainer(containerID string) error
	StopContainer(containerID string, timeout *int) error
	RestartContainer(containerID string, timeout *int) error
	RemoveContainer(containerID string, force bool) error
	ContainerLogs(containerID string, follow bool) (io.ReadCloser, error)
	ExecInContainer(containerID string, cmd []string) (string, error)
	Stats(containerID string, stream bool) (io.ReadCloser, error)
	PruneContainers() (types.ContainersPruneReport, error)
}

type containerService struct {
	cli *client.Client
}

// NewContainerService creates a new instance of ContainerService using the Docker SDK
func NewContainerService() (ContainerService, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &containerService{cli: cli}, nil
}

func (s *containerService) ListContainers(all bool) ([]types.Container, error) {
	containers, err := s.cli.ContainerList(context.Background(), container.ListOptions{
		All: all,
	})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (s *containerService) InspectContainer(containerID string) (*types.ContainerJSON, error) {
	containerJSON, err := s.cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect container: %w", err)
	}
	return &containerJSON, nil
}

func (s *containerService) StartContainer(containerID string) error {
	if err := s.cli.ContainerStart(context.Background(), containerID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}
	log.Printf("Container %s started successfully", containerID)
	return nil
}

func (s *containerService) StopContainer(containerID string, timeout *int) error {
	if err := s.cli.ContainerStop(context.Background(), containerID, container.StopOptions{Timeout: timeout}); err != nil {
		return fmt.Errorf("failed to stop container: %w", err)
	}
	log.Printf("Container %s stopped successfully", containerID)
	return nil
}

func (s *containerService) RestartContainer(containerID string, timeout *int) error {
	if err := s.cli.ContainerRestart(context.Background(), containerID, container.StopOptions{Timeout: timeout}); err != nil {
		return fmt.Errorf("failed to restart container: %w", err)
	}
	log.Printf("Container %s restarted successfully", containerID)
	return nil
}

func (s *containerService) RemoveContainer(containerID string, force bool) error {
	options := container.RemoveOptions{Force: force}
	if err := s.cli.ContainerRemove(context.Background(), containerID, options); err != nil {
		return fmt.Errorf("failed to remove container: %w", err)
	}
	log.Printf("Container %s removed successfully", containerID)
	return nil
}

func (s *containerService) ContainerLogs(containerID string, follow bool) (io.ReadCloser, error) {
	logs, err := s.cli.ContainerLogs(context.Background(), containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     follow,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs for container %s: %w", containerID, err)
	}
	return logs, nil
}

func (s *containerService) ExecInContainer(containerID string, cmd []string) (string, error) {
	execResp, err := s.cli.ContainerExecCreate(context.Background(), containerID, types.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create exec instance: %w", err)
	}

	attachResp, err := s.cli.ContainerExecAttach(context.Background(), execResp.ID, types.ExecStartCheck{})
	if err != nil {
		return "", fmt.Errorf("failed to attach to exec instance: %w", err)
	}
	defer attachResp.Close()

	output, err := io.ReadAll(attachResp.Reader)
	if err != nil {
		return "", fmt.Errorf("failed to read exec output: %w", err)
	}

	return string(output), nil
}

func (s *containerService) Stats(containerID string, stream bool) (io.ReadCloser, error) {
	stats, err := s.cli.ContainerStats(context.Background(), containerID, stream)
	if err != nil {
		return nil, fmt.Errorf("failed to get stats for container %s: %w", containerID, err)
	}
	return stats.Body, nil
}

func (s *containerService) PruneContainers() (types.ContainersPruneReport, error) {
	report, err := s.cli.ContainersPrune(context.Background(), filters.Args{})
	if err != nil {
		return types.ContainersPruneReport{}, fmt.Errorf("failed to prune containers: %w", err)
	}
	return report, nil
}
