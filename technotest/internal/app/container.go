package app

import (
	"context"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technotest/config"
	"technotest/internal/consts"
	"technotest/internal/repository"
	"technotest/internal/service"
	"technotest/pkg/postgres"
)

type (
	Container struct {
		Services Services
		servers  servers
	}

	Services struct {
		ContainerService  *service.ContainerService
		ContainerExplorer *service.ContainerExplorer
	}

	repositories struct {
		ContainerRepo *repository.ContainerRepository
	}

	clients struct {
		dockerClient *docker.Client
	}

	servers struct {
		httpRouter *mux.Router
		httpServer *http.Server
	}
)

func NewContainer(ctx context.Context, cfg *config.Config) *Container {
	sClients := initClients(cfg)
	sRepositories := initRepos(ctx, cfg)

	return &Container{
		Services: initServices(sClients, sRepositories),
		servers:  initServers(cfg),
	}
}

func initClients(cfg *config.Config) clients {
	dockerClient, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal("Failed to init docker client", err)
	}

	return clients{
		dockerClient: dockerClient,
	}
}

func initServices(sClients clients, sRepos repositories) Services {
	return Services{
		ContainerService:  service.NewContainerService(sRepos.ContainerRepo),
		ContainerExplorer: service.NewContainerExplorer(sClients.dockerClient),
	}
}

func initRepos(ctx context.Context, cfg *config.Config) repositories {
	pgClient, err := postgres.NewClient(ctx, cfg.URL)
	if err != nil {
		log.Fatal("init pg client", err)
	}

	return repositories{
		ContainerRepo: repository.NewContainerRepository(pgClient),
	}
}

func initServers(cfg *config.Config) servers {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    consts.Localhost + ":" + cfg.Port,
		Handler: router,
	}

	return servers{
		httpRouter: router,
		httpServer: server,
	}
}
