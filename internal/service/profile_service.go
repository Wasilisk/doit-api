package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/sqlc"
)

type ProfileService struct {
	profileRepo *repository.ProfileRepository
}

type UpdateProfileInput struct {
	UserID    uuid.UUID
	FullName  *string
	AvatarURL *string
}

const avatarsDir = "static/avatars"

func NewProfileService(profileRepo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{profileRepo: profileRepo}
}

func (s *ProfileService) GetProfile(ctx context.Context, userID uuid.UUID) (dto.ProfileResponse, error) {
	profile, err := s.profileRepo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return dto.ProfileResponse{}, err
	}
	return toProfileResponse(profile), nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, input UpdateProfileInput) (dto.ProfileResponse, error) {
	profile, err := s.profileRepo.UpdateProfile(ctx, repository.UpdateProfileInput{
		UserID:    input.UserID,
		FullName:  input.FullName,
		AvatarURL: input.AvatarURL,
	})
	if err != nil {
		return dto.ProfileResponse{}, err
	}
	return toProfileResponse(profile), nil
}

func (s *ProfileService) UploadAvatar(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(header.Filename)
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowed[ext] {
		return "", fmt.Errorf("file type %s not allowed", ext)
	}

	if err := os.MkdirAll(avatarsDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create avatars dir: %w", err)
	}

	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	destPath := filepath.Join(avatarsDir, filename)

	dest, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return fmt.Sprintf("/static/avatars/%s", filename), nil
}

func toProfileResponse(p sqlc.UserProfile) dto.ProfileResponse {
	var avatar *string
	if p.AvatarUrl.Valid {
		avatar = &p.AvatarUrl.String
	}

	return dto.ProfileResponse{
		ID:        p.ID.String(),
		UserID:    p.UserID.String(),
		FullName:  p.FullName,
		AvatarURL: avatar,
	}
}
