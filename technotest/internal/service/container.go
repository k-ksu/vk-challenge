package service

import (
	"context"

	"technotest/internal/entity"
)

type (
	ContainerRepository interface {
		UpdateContainersStatus(
			ctx context.Context,
			containers []entity.ContainerStatus,
		) error
		GetContainersStatus(ctx context.Context) ([]entity.ContainerStatus, error)
	}
	ContainerService struct {
		repo ContainerRepository
	}
)

func NewContainerService(repo ContainerRepository) *ContainerService {
	return &ContainerService{
		repo: repo,
	}
}

func (cs *ContainerService) UpdateContainersStatus(
	ctx context.Context,
	containers []entity.ContainerStatus,
) error {
	return cs.repo.UpdateContainersStatus(ctx, containers)
}

func (cs *ContainerService) GetContainersStatus(ctx context.Context) ([]entity.ContainerStatus, error) {
	return cs.repo.GetContainersStatus(ctx)
}
