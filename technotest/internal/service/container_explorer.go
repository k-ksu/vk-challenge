package service

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
	"technotest/internal/entity"
)

type ContainerExplorer struct {
	client *docker.Client
}

func NewContainerExplorer(client *docker.Client) *ContainerExplorer {
	return &ContainerExplorer{
		client: client,
	}
}

func (ce *ContainerExplorer) ListAllContainersStatus() ([]entity.ContainerStatus, error) {
	conts, err := ce.client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("client.ListContainers: %w", err)
	}

	containersStatus := make([]entity.ContainerStatus, 0, len(conts))

	for _, cont := range conts {
		name := "undefined"
		if len(cont.Names) > 0 {
			name = cont.Names[0]
		}

		containersStatus = append(containersStatus, entity.ContainerStatus{
			ID:     cont.ID,
			Name:   name,
			Image:  cont.Image,
			State:  cont.State,
			Status: cont.Status,
		})
	}

	return containersStatus, nil
}
