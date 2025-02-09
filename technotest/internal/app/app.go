package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"technotest/config"
	httpcontroller "technotest/internal/controller/http"
)

func Run(ctx context.Context, cfg *config.Config) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	cont := NewContainer(ctx, cfg)

	impl := httpcontroller.NewTechPointAPI(cont.Services.ContainerService)
	impl.RegisterGateway(cont.servers.httpRouter)

	go func() {
		if err := cont.servers.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, "Server error")
		}
	}()

	log.Println("Server started at address:", cont.servers.httpServer.Addr)

	<-quit
	log.Println("Shutting down server...")

	if err := cont.servers.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	return nil
}
