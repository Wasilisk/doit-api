package service

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/sqlc"
	"github.com/wasilisk/doit-api/internal/storage"
)

type ProfileService struct {
	profileRepo   *repository.ProfileRepository
	avatarStorage *storage.AvatarStorage
}

type UpdateProfileInput struct {
	UserID    uuid.UUID
	FullName  *string
	AvatarURL *string
}

func NewProfileService(profileRepo *repository.ProfileRepository, avatarStorage *storage.AvatarStorage) *ProfileService {
	return &ProfileService{profileRepo: profileRepo, avatarStorage: avatarStorage}
}

func (s *ProfileService) GetProfile(ctx context.Context, userID uuid.UUID) (dto.ProfileResponse, error) {
	profile, err := s.profileRepo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return dto.ProfileResponse{}, apperror.New(apperror.CodeProfileNotFound)
	}
	return toProfileResponse(profile), nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, input UpdateProfileInput) (dto.ProfileResponse, error) {
	existing, err := s.profileRepo.GetProfileByUserID(ctx, input.UserID)
	if err != nil {
		return dto.ProfileResponse{}, apperror.New(apperror.CodeProfileNotFound)
	}

	updated, err := s.profileRepo.UpdateProfile(ctx, repository.UpdateProfileInput{
		UserID:    input.UserID,
		FullName:  input.FullName,
		AvatarURL: input.AvatarURL,
	})
	if err != nil {
		return dto.ProfileResponse{}, apperror.New(apperror.CodeInternal)
	}

	if input.AvatarURL != nil && existing.AvatarUrl.Valid {
		if err := s.avatarStorage.Delete(existing.AvatarUrl.String); err != nil {
			return dto.ProfileResponse{}, apperror.New(apperror.CodeAvatarUploadFailed)
		}
	}

	return toProfileResponse(sqlc.GetProfileByUserIDRow{
		ID:        existing.ID,
		UserID:    existing.UserID,
		Email:     existing.Email,
		FullName:  updated.FullName,
		AvatarUrl: updated.AvatarUrl,
	}), nil
}

func (s *ProfileService) UploadAvatar(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (string, error) {
	if err := s.avatarStorage.Validate(header.Filename); err != nil {
		return "", err
	}

	ext := filepath.Ext(header.Filename)
	filename, err := s.avatarStorage.Save(file, ext)
	if err != nil {
		return "", err
	}

	return s.avatarStorage.PublicURL(filename), nil
}

func toProfileResponse(p sqlc.GetProfileByUserIDRow) dto.ProfileResponse {
	var avatar *string
	if p.AvatarUrl.Valid {
		avatar = &p.AvatarUrl.String
	}

	return dto.ProfileResponse{
		ID:        p.ID.String(),
		UserID:    p.UserID.String(),
		Email:     p.Email,
		FullName:  p.FullName,
		AvatarURL: avatar,
	}
}
