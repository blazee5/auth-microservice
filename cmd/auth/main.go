package main

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/app"
	"github.com/blazee5/auth-microservice/internal/config"
	"github.com/blazee5/auth-microservice/internal/repository"
	"github.com/blazee5/auth-microservice/internal/service"
	postgresLib "github.com/blazee5/auth-microservice/lib/db/postgres"
	"github.com/blazee5/auth-microservice/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	db := postgresLib.Run(ctx, cfg)
	repo := repository.NewRepository(db)
	services := service.NewService(log, repo)
	server := app.New(cfg, log, services)

	log.Log.Infof("server starting on port %s...", cfg.Server.Port)

	go func() {
		server.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.Stop()

	cancel()
}
