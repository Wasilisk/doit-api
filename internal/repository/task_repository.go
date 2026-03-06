package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/sqlc"

	utils "github.com/wasilisk/doit-api/internal/utils"
)

type CreateTaskInput struct {
	UserID      uuid.UUID
	Name        string
	Description *string
	Date        *time.Time
	TimeStart   *time.Time
	TimeEnd     *time.Time
}

type UpdateTaskInput struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        *string
	Description *string
	Date        *time.Time
	TimeStart   *time.Time
	TimeEnd     *time.Time
	IsCompleted *bool
	IsFavourite *bool
}

type TaskFilterInput struct {
	UserID      uuid.UUID
	Date        *time.Time
	TagID       *string
	IsCompleted *bool
	IsDeleted   *bool
}

type TaskRepository struct {
	queries *sqlc.Queries
	db      *sql.DB
}

func NewTaskRepository(database *sql.DB) *TaskRepository {
	return &TaskRepository{queries: sqlc.New(database), db: database}
}

func (r *TaskRepository) CreateTask(ctx context.Context, input CreateTaskInput) (sqlc.Task, error) {
	params := sqlc.CreateTaskParams{
		UserID:      input.UserID,
		Name:        input.Name,
		Description: utils.StringToNullString(input.Description),
		Date:        utils.NullTimeFrom(input.Date),
		TimeStart:   utils.NullTimeFrom(input.TimeStart),
		TimeEnd:     utils.NullTimeFrom(input.TimeEnd),
	}
	return r.queries.CreateTask(ctx, params)
}

func (r *TaskRepository) UpdateTask(ctx context.Context, input UpdateTaskInput) (sqlc.Task, error) {
	existing, err := r.queries.GetTaskByID(ctx, sqlc.GetTaskByIDParams{ID: input.ID, UserID: input.UserID})
	if err != nil {
		return sqlc.Task{}, err
	}

	name := existing.Name
	if input.Name != nil {
		name = *input.Name
	}

	description := existing.Description
	if input.Description != nil {
		description = sql.NullString{String: *input.Description, Valid: true}
	}

	date := existing.Date
	if input.Date != nil {
		date = utils.NullTimeFrom(input.Date)
	}

	timeStart := existing.TimeStart
	if input.TimeStart != nil {
		timeStart = utils.NullTimeFrom(input.TimeStart)
	}

	timeEnd := existing.TimeEnd
	if input.TimeEnd != nil {
		timeEnd = utils.NullTimeFrom(input.TimeEnd)
	}

	isCompleted := existing.IsCompleted
	if input.IsCompleted != nil {
		isCompleted = *input.IsCompleted
	}

	isFavourite := existing.IsFavourite
	if input.IsFavourite != nil {
		isFavourite = *input.IsFavourite
	}

	return r.queries.UpdateTask(ctx, sqlc.UpdateTaskParams{
		ID:          input.ID,
		UserID:      input.UserID,
		Name:        name,
		Description: description,
		Date:        date,
		TimeStart:   timeStart,
		TimeEnd:     timeEnd,
		IsCompleted: isCompleted,
		IsFavourite: isFavourite,
	})
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id, userID uuid.UUID) (sqlc.Task, error) {
	return r.queries.GetTaskByID(ctx, sqlc.GetTaskByIDParams{ID: id, UserID: userID})
}

func (r *TaskRepository) GetTasks(ctx context.Context, filter TaskFilterInput) ([]sqlc.Task, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psql.Select(
		"tasks.id",
		"tasks.user_id",
		"tasks.name",
		"tasks.description",
		"tasks.date",
		"tasks.time_start",
		"tasks.time_end",
		"tasks.is_completed",
		"tasks.is_favourite",
		"tasks.deleted_at",
		"tasks.created_at",
		"tasks.updated_at",
	).From("tasks").
		Where(sq.Eq{"user_id": filter.UserID})

	if filter.IsDeleted != nil && *filter.IsDeleted {
		qb = qb.Where(sq.NotEq{"deleted_at": nil})
	} else {
		qb = qb.Where(sq.Eq{"deleted_at": nil})
	}

	if filter.Date != nil {
		t := filter.Date.UTC()
		dayStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		dayEnd := dayStart.Add(24 * time.Hour)
		qb = qb.Where(sq.GtOrEq{"date": dayStart}).Where(sq.Lt{"date": dayEnd})
	}

	if filter.IsCompleted != nil {
		qb = qb.Where(sq.Eq{"is_completed": *filter.IsCompleted})
	}

	if filter.TagID != nil {
		qb = qb.
			Join("task_tags ON task_tags.task_id = tasks.id").
			Where(sq.Eq{"task_tags.tag_id": *filter.TagID})
	}

	qb = qb.OrderBy("created_at DESC")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []sqlc.Task
	for rows.Next() {
		var t sqlc.Task
		if err := rows.Scan(
			&t.ID, &t.UserID, &t.Name, &t.Description,
			&t.Date, &t.TimeStart, &t.TimeEnd,
			&t.IsCompleted, &t.IsFavourite,
			&t.DeletedAt, &t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) SoftDeleteTask(ctx context.Context, id, userID uuid.UUID) error {
	return r.queries.SoftDeleteTask(ctx, sqlc.SoftDeleteTaskParams{ID: id, UserID: userID})
}

func (r *TaskRepository) RestoreTask(ctx context.Context, id, userID uuid.UUID) error {
	return r.queries.RestoreTask(ctx, sqlc.RestoreTaskParams{ID: id, UserID: userID})
}

func (r *TaskRepository) AddTagToTask(ctx context.Context, taskID, tagID uuid.UUID) error {
	return r.queries.AddTagToTask(ctx, sqlc.AddTagToTaskParams{TaskID: taskID, TagID: tagID})
}

func (r *TaskRepository) RemoveTagFromTask(ctx context.Context, taskID, tagID uuid.UUID) error {
	return r.queries.RemoveTagFromTask(ctx, sqlc.RemoveTagFromTaskParams{TaskID: taskID, TagID: tagID})
}

func (r *TaskRepository) GetTagsByTaskID(ctx context.Context, taskID uuid.UUID) ([]sqlc.Tag, error) {
	return r.queries.GetTagsByTaskID(ctx, taskID)
}
