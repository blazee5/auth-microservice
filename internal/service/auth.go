package service

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/repository"
	"github.com/blazee5/auth-microservice/lib/auth"
	"github.com/blazee5/auth-microservice/lib/logger"
	pb "github.com/blazee5/protos/auth"
)

type AuthService struct {
	log  *logger.Logger
	repo *repository.Repository
}

func NewAuthService(log *logger.Logger, repo *repository.Repository) *AuthService {
	return &AuthService{log: log, repo: repo}
}

func (s *AuthService) SignUp(ctx context.Context, in *pb.SignUpRequest) (string, error) {
	in.Password = auth.GenerateHashPassword(in.Password)
	return s.repo.CreateUser(ctx, in)
}

func (s *AuthService) SignIn(ctx context.Context, in *pb.SignInRequest) (auth.Tokens, error) {
	in.Password = auth.GenerateHashPassword(in.Password)
	user, err := s.repo.ValidateUser(ctx, in.Email, in.Password)

	if err != nil {
		return auth.Tokens{}, err
	}

	return auth.GenerateTokens(user.ID)
}

func (s *AuthService) RefreshTokens(ctx context.Context, input *pb.TokenRequest) (auth.Tokens, error) {
	userID, err := auth.ParseToken(input.RefreshToken)

	if err != nil {
		return auth.Tokens{}, err
	}

	return auth.GenerateTokens(userID)
}
