package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/sqlc"

	utils "github.com/wasilisk/doit-api/internal/utils"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(ctx context.Context, userID uuid.UUID, req dto.CreateTaskRequest) (dto.TaskResponse, error) {
	task, err := s.taskRepo.CreateTask(ctx, repository.CreateTaskInput{
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
		Date:        req.Date,
		TimeStart:   req.TimeStart,
		TimeEnd:     req.TimeEnd,
	})
	if err != nil {
		return dto.TaskResponse{}, err
	}

	for _, tagID := range req.TagIDs {
		tid, err := uuid.Parse(tagID)
		if err != nil {
			continue
		}
		_ = s.taskRepo.AddTagToTask(ctx, task.ID, tid)
	}

	return s.buildTaskResponse(ctx, task)
}

func (s *TaskService) UpdateTask(ctx context.Context, userID, taskID uuid.UUID, req dto.UpdateTaskRequest) (dto.TaskResponse, error) {
	task, err := s.taskRepo.UpdateTask(ctx, repository.UpdateTaskInput{
		ID:          taskID,
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
		Date:        req.Date,
		TimeStart:   req.TimeStart,
		TimeEnd:     req.TimeEnd,
		IsCompleted: req.IsCompleted,
		IsFavourite: req.IsFavourite,
	})
	if err != nil {
		return dto.TaskResponse{}, errors.New("task not found")
	}

	if req.TagIDs != nil {
		existing, _ := s.taskRepo.GetTagsByTaskID(ctx, taskID)
		existingMap := make(map[string]bool)
		for _, t := range existing {
			existingMap[t.ID.String()] = true
		}

		newMap := make(map[string]bool)
		for _, id := range req.TagIDs {
			newMap[id] = true
		}

		for _, id := range req.TagIDs {
			if !existingMap[id] {
				if tid, err := uuid.Parse(id); err == nil {
					_ = s.taskRepo.AddTagToTask(ctx, taskID, tid)
				}
			}
		}

		for _, t := range existing {
			if !newMap[t.ID.String()] {
				_ = s.taskRepo.RemoveTagFromTask(ctx, taskID, t.ID)
			}
		}
	}

	return s.buildTaskResponse(ctx, task)
}

func (s *TaskService) GetTasks(ctx context.Context, userID uuid.UUID, filter dto.TaskFilterRequest) ([]dto.TaskResponse, error) {
	tasks, err := s.taskRepo.GetTasks(ctx, repository.TaskFilterInput{
		UserID:      userID,
		Date:        filter.Date,
		TagID:       filter.TagID,
		IsCompleted: filter.IsCompleted,
		IsDeleted:   filter.IsDeleted,
	})
	if err != nil {
		return nil, err
	}

	result := make([]dto.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		resp, err := s.buildTaskResponse(ctx, task)
		if err != nil {
			continue
		}
		result = append(result, resp)
	}
	return result, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, userID, taskID uuid.UUID) (dto.TaskResponse, error) {
	task, err := s.taskRepo.GetTaskByID(ctx, taskID, userID)
	if err != nil {
		return dto.TaskResponse{}, errors.New("task not found")
	}
	return s.buildTaskResponse(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, userID, taskID uuid.UUID) error {
	_, err := s.taskRepo.GetTaskByID(ctx, taskID, userID)
	if err != nil {
		return errors.New("task not found")
	}
	return s.taskRepo.SoftDeleteTask(ctx, taskID, userID)
}

func (s *TaskService) RestoreTask(ctx context.Context, userID, taskID uuid.UUID) error {
	return s.taskRepo.RestoreTask(ctx, taskID, userID)
}

func (s *TaskService) buildTaskResponse(ctx context.Context, task sqlc.Task) (dto.TaskResponse, error) {
	tags, _ := s.taskRepo.GetTagsByTaskID(ctx, task.ID)

	tagResponses := make([]dto.TagResponse, len(tags))
	for i, tag := range tags {
		tagResponses[i] = dto.TagResponse{
			ID:     tag.ID.String(),
			UserID: tag.UserID.String(),
			Name:   tag.Name,
			Color:  tag.Color,
		}
	}

	return dto.TaskResponse{
		ID:          task.ID.String(),
		UserID:      task.UserID.String(),
		Name:        task.Name,
		IsCompleted: task.IsCompleted,
		IsFavourite: task.IsFavourite,
		Tags:        tagResponses,
		CreatedAt:   task.CreatedAt.Time.Format("2006-01-02T15:04:05Z"),
		Description: utils.NullStringToPtr(task.Description),
		Date:        utils.NullTimeToUnix(task.Date),
		TimeStart:   utils.NullTimeToUnix(task.TimeStart),
		TimeEnd:     utils.NullTimeToUnix(task.TimeEnd),
		DeletedAt:   utils.NullTimeToStringPtr(task.DeletedAt),
	}, nil
}
