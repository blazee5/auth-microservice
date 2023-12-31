package repository

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/models"
	"github.com/blazee5/auth-microservice/internal/repository/postgres"
	pb "github.com/blazee5/protos/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	PostgresRepository
}

type PostgresRepository interface {
	CreateUser(ctx context.Context, in *pb.SignUpRequest) (string, error)
	ValidateUser(ctx context.Context, email, password string) (models.User, error)
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		PostgresRepository: postgres.NewAuthRepository(db),
	}
}
