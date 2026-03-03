package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{queries: sqlc.New(database)}
}

func (r *UserRepository) CreateUser(ctx context.Context, email, hashedPassword string) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
	})
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, id)
}
