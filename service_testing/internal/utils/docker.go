package utils

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetDockerContainerInfo(ctx context.Context, containerName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("error creating Docker client: %w", err)
	}

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return fmt.Errorf("error listing containers: %w", err)
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return nil
			}
		}
	}

	return fmt.Errorf("container with name %s not found", containerName)
}
