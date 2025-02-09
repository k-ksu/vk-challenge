package main

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"
	"technotest/config"
	"technotest/internal/app"
	"technotest/internal/client/techotestapi"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	cont := app.NewContainer(ctx, cfg)

	cl := techotestapi.NewTechnoTestAPIClient(cfg.TechApiClientHost, cont.Services.ContainerExplorer)

	c := cron.New()
	if _, err = c.AddFunc("@every "+cfg.Interval, func() {
		if cronErr := cl.UpdateContainersStatus(); cronErr != nil {
			log.Println("Failed to ping", cronErr)
		}
	}); err != nil {
		log.Fatal("Adding new func to cron failed", err)
	}

	c.Run()
}
