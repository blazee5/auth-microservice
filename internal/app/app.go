package app

import (
	"context"
	"fmt"
	"github.com/blazee5/auth-microservice/internal/config"
	authgrpc "github.com/blazee5/auth-microservice/internal/grpc/auth"
	"github.com/blazee5/auth-microservice/internal/service"
	"github.com/blazee5/auth-microservice/lib/logger"
	pb "github.com/blazee5/protos/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
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

	go func() {
		if err := a.grpcServer.Serve(l); err != nil {
			log.Fatalf("error while serve grpc server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:"+a.cfg.Server.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	a.log.Log.Infof("Serving gRPC-Gateway on http://0.0.0.0:8080")
	a.log.Log.Fatalln(gwServer.ListenAndServe())
}

func (a *App) Stop() {
	a.grpcServer.GracefulStop()
}
