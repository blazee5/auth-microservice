package app

import (
	"fmt"
	"github.com/blazee5/auth-microservice/internal/config"
	authgrpc "github.com/blazee5/auth-microservice/internal/grpc/auth"
	"github.com/blazee5/auth-microservice/internal/service"
	"github.com/blazee5/auth-microservice/lib/logger"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	cfg        *config.Config
	log        *logger.Logger
	grpcServer *grpc.Server
}

func New(cfg *config.Config, log *logger.Logger, service *service.Service) *App {
	server := grpc.NewServer()

	authgrpc.Register(log, service, server)

	return &App{
		cfg:        cfg,
		log:        log,
		grpcServer: server,
	}
}

func (a *App) Run() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", a.cfg.Server.Port))

	if err != nil {
		log.Fatalf("error while listen grpc server: %v", err)
	}

	if err := a.grpcServer.Serve(l); err != nil {
		log.Fatalf("error while serve grpc server: %v", err)
	}
}

func (a *App) Stop() {
	a.grpcServer.GracefulStop()
}
