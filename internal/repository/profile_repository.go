package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/sqlc"
	dbutils "github.com/wasilisk/doit-api/internal/utils/db"
)

type ProfileRepository struct {
	queries *sqlc.Queries
}

type CreateProfileInput struct {
	UserID    uuid.UUID
	FullName  string
	AvatarURL *string
}

type UpdateProfileInput struct {
	UserID    uuid.UUID
	FullName  *string
	AvatarURL *string
}

func NewProfileRepository(database *sql.DB) *ProfileRepository {
	return &ProfileRepository{queries: sqlc.New(database)}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, input CreateProfileInput) (sqlc.UserProfile, error) {
	return r.queries.CreateProfile(ctx, sqlc.CreateProfileParams{
		UserID:    input.UserID,
		FullName:  input.FullName,
		AvatarUrl: dbutils.NullString(input.AvatarURL),
	})
}

func (r *ProfileRepository) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (sqlc.GetProfileByUserIDRow, error) {
	return r.queries.GetProfileByUserID(ctx, userID)
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, input UpdateProfileInput) (sqlc.UserProfile, error) {
	return r.queries.UpdateProfile(ctx, sqlc.UpdateProfileParams{
		UserID:    input.UserID,
		FullName:  dbutils.NullString(input.FullName),
		AvatarUrl: dbutils.NullString(input.AvatarURL),
	})
}
