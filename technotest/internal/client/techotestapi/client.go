package techotestapi

import (
	"technotest/internal/entity"
)

type (
	ContainerExplorer interface {
		ListAllContainersStatus() ([]entity.ContainerStatus, error)
	}

	TechTestAPIClient struct {
		Addr              string
		ContainerExplorer ContainerExplorer
	}
)

func NewTechnoTestAPIClient(Addr string, Explorer ContainerExplorer) *TechTestAPIClient {
	return &TechTestAPIClient{
		Addr:              Addr,
		ContainerExplorer: Explorer,
	}
}
