package auth

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/service"
	"github.com/blazee5/auth-microservice/lib/logger"
	pb "github.com/blazee5/protos/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	log     *logger.Logger
	service *service.Service
	pb.UnimplementedAuthServer
}

func Register(log *logger.Logger, service *service.Service, server *grpc.Server) {
	pb.RegisterAuthServer(server, &Server{log: log, service: service})
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	id, err := s.service.SignUp(ctx, in)

	if err != nil {
		s.log.Log.Infof("error while sign up: %v", err)

		return nil, status.Error(codes.Internal, "server error")
	}

	return &pb.SignUpResponse{UserId: id}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	token, err := s.service.SignIn(ctx, in)

	if err != nil {
		s.log.Log.Infof("error while sign in: %v", err)

		return nil, status.Error(codes.Internal, "server error")
	}

	return &pb.SignInResponse{Token: token}, nil
}
