package postgres

import (
	"context"
	"github.com/blazee5/auth-microservice/internal/models"
	pb "github.com/blazee5/protos/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo *AuthRepository) CreateUser(ctx context.Context, in *pb.SignUpRequest) (string, error) {
	var id string

	err := repo.db.QueryRow(ctx, "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		in.FirstName, in.LastName, in.Email, in.Password).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (repo *AuthRepository) ValidateUser(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	err := repo.db.QueryRow(ctx, "SELECT id, first_name, last_name, email FROM users WHERE email = $1 AND password = $2",
		email, password).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
