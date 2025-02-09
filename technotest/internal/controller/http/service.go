package http

import (
	"context"

	"github.com/gorilla/mux"
	"technotest/internal/entity"
)

type (
	ContainerService interface {
		UpdateContainersStatus(
			ctx context.Context,
			containers []entity.ContainerStatus,
		) error
		GetContainersStatus(ctx context.Context) ([]entity.ContainerStatus, error)
	}
	TechPointAPI struct {
		containerService ContainerService
	}
)

func NewTechPointAPI(containerService ContainerService) *TechPointAPI {
	return &TechPointAPI{
		containerService: containerService,
	}
}

func (t *TechPointAPI) RegisterGateway(mux *mux.Router) {
	mux.Handle("/update_containers_status", t.UpdateContainersStatus()).Methods("POST")
	mux.Handle("/get_containers_status", t.GetContainersStatus()).Methods("GET")
}
