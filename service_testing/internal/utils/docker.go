package utils

import (
	"context"
	"errors"
	internalErr "service_testing/internal/errors"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetDockerContainerInfo(ctx context.Context, containerName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return errors.Join(internalErr.ErrDockerClientCreation, err)
	}

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return errors.Join(internalErr.ErrDockerGetAllContainer, err)
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return nil
			}
		}
	}

	return internalErr.FailDockerContainerNotFound
}
