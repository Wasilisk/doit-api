package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/sqlc"
)

type ProfileService struct {
	profileRepo *repository.ProfileRepository
}

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

func (s *ProfileService) UpdateProfile(ctx context.Context, userID uuid.UUID, req *dto.UpdateProfileRequest) (dto.ProfileResponse, error) {
	profile, err := s.profileRepo.UpdateProfile(ctx, repository.UpdateProfileInput{
		UserID:    userID,
		FullName:  &req.FullName,
		AvatarURL: &req.AvatarURL,
	})
	if err != nil {
		return dto.ProfileResponse{}, err
	}
	return toProfileResponse(profile), nil
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
