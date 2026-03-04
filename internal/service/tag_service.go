package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/sqlc"
)

type TagService struct {
	tagRepo *repository.TagRepository
}

func NewTagService(tagRepo *repository.TagRepository) *TagService {
	return &TagService{tagRepo: tagRepo}
}

func (s *TagService) CreateTag(ctx context.Context, userID uuid.UUID, req dto.CreateTagRequest) (dto.TagResponse, error) {
	tag, err := s.tagRepo.CreateTag(ctx, repository.CreateTagInput{
		UserID: userID,
		Name:   req.Name,
		Color:  req.Color,
	})
	if err != nil {
		return dto.TagResponse{}, errors.New("tag with this name already exists")
	}
	return toTagResponse(tag), nil
}

func (s *TagService) GetTags(ctx context.Context, userID uuid.UUID) ([]dto.TagResponse, error) {
	tags, err := s.tagRepo.GetTagsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.TagResponse, len(tags))
	for i, tag := range tags {
		result[i] = toTagResponse(tag)
	}
	return result, nil
}

func (s *TagService) UpdateTag(ctx context.Context, userID uuid.UUID, tagID uuid.UUID, req dto.UpdateTagRequest) (dto.TagResponse, error) {
	existing, err := s.tagRepo.GetTagByID(ctx, tagID, userID)
	if err != nil {
		return dto.TagResponse{}, errors.New("tag not found")
	}

	name := existing.Name
	color := existing.Color
	if req.Name != nil {
		name = *req.Name
	}
	if req.Color != nil {
		color = *req.Color
	}

	tag, err := s.tagRepo.UpdateTag(ctx, repository.UpdateTagInput{
		ID:     tagID,
		UserID: userID,
		Name:   name,
		Color:  color,
	})
	if err != nil {
		return dto.TagResponse{}, err
	}
	return toTagResponse(tag), nil
}

func (s *TagService) DeleteTag(ctx context.Context, userID, tagID uuid.UUID) error {
	_, err := s.tagRepo.GetTagByID(ctx, tagID, userID)
	if err != nil {
		return errors.New("tag not found")
	}
	return s.tagRepo.DeleteTag(ctx, tagID, userID)
}

func toTagResponse(t sqlc.Tag) dto.TagResponse {
	return dto.TagResponse{
		ID:        t.ID.String(),
		UserID:    t.UserID.String(),
		Name:      t.Name,
		Color:     t.Color,
		CreatedAt: t.CreatedAt.Time.Format("2006-01-02T15:04:05Z"),
	}
}
