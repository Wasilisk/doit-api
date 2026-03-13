package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/sqlc"
)

type CreateTagInput struct {
	UserID uuid.UUID
	Name   string
	Color  string
}

type UpdateTagInput struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Name   string
	Color  string
}

type TagRepository struct {
	queries *sqlc.Queries
}

func NewTagRepository(database *sql.DB) *TagRepository {
	return &TagRepository{queries: sqlc.New(database)}
}

func (r *TagRepository) CreateTag(ctx context.Context, input CreateTagInput) (sqlc.Tag, error) {
	return r.queries.CreateTag(ctx, sqlc.CreateTagParams{
		UserID: input.UserID,
		Name:   input.Name,
		Color:  input.Color,
	})
}

func (r *TagRepository) GetTagByID(ctx context.Context, id, userID uuid.UUID) (sqlc.Tag, error) {
	return r.queries.GetTagByID(ctx, sqlc.GetTagByIDParams{
		ID:     id,
		UserID: userID,
	})
}

func (r *TagRepository) GetTagsByUserID(ctx context.Context, userID uuid.UUID) ([]sqlc.GetTagsByUserIDRow, error) {
	return r.queries.GetTagsByUserID(ctx, userID)
}

func (r *TagRepository) UpdateTag(ctx context.Context, input UpdateTagInput) (sqlc.Tag, error) {
	return r.queries.UpdateTag(ctx, sqlc.UpdateTagParams{
		ID:     input.ID,
		UserID: input.UserID,
		Name:   input.Name,
		Color:  input.Color,
	})
}

func (r *TagRepository) DeleteTag(ctx context.Context, id, userID uuid.UUID) error {
	return r.queries.DeleteTag(ctx, sqlc.DeleteTagParams{
		ID:     id,
		UserID: userID,
	})
}
