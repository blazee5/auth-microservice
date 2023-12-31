package service

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/repository"
	"github.com/blazee5/auth-microservice/lib/logger"
	pb "github.com/blazee5/protos/auth"
)

type Service struct {
	Auth
}

type Auth interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error)
	SignIn(ctx context.Context, input *pb.SignInRequest) (string, error)
}

func NewService(log *logger.Logger, repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(log, repo),
	}
}
