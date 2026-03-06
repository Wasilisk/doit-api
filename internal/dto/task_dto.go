package dto

import "time"

type CreateTaskRequest struct {
	Name        string   `json:"name"        binding:"required"`
	Description *string  `json:"description"`
	Date        *int64   `json:"date"`
	TimeStart   *int64   `json:"time_start"`
	TimeEnd     *int64   `json:"time_end"`
	TagIDs      []string `json:"tag_ids"`
}

type UpdateTaskRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Date        *int64   `json:"date"`
	TimeStart   *int64   `json:"time_start"`
	TimeEnd     *int64   `json:"time_end"`
	TagIDs      []string `json:"tag_ids"`
	IsCompleted *bool    `json:"is_completed"`
	IsFavourite *bool    `json:"is_favourite"`
}

type TaskFilterRequest struct {
	Date        *int64  `form:"date"`
	TagID       *string `form:"tag_id"`
	IsCompleted *bool   `form:"is_completed"`
	IsDeleted   *bool   `form:"is_deleted"`
}

type TaskResponse struct {
	ID          string        `json:"id"`
	UserID      string        `json:"user_id"`
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	Date        *int64        `json:"date"`
	TimeStart   *int64        `json:"time_start"`
	TimeEnd     *int64        `json:"time_end"`
	IsCompleted bool          `json:"is_completed"`
	IsFavourite bool          `json:"is_favourite"`
	DeletedAt   *time.Time    `json:"deleted_at"`
	Tags        []TagResponse `json:"tags"`
	CreatedAt   time.Time     `json:"created_at"`
}
