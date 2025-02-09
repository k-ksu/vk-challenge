package main

import (
	"context"
	"log"

	"technotest/config"
	"technotest/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	if err = app.Run(ctx, cfg); err != nil {
		log.Fatal("App execution failed", err)
	}
}
