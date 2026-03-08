package service

import (
	"context"

	apperror "github.com/wasilisk/doit-api/internal/app_error"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/utils"
	dbutils "github.com/wasilisk/doit-api/internal/utils/db"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	profileRepo *repository.ProfileRepository
	jwtSecret   string
}

type RegisterInput struct {
	Email    string
	Password string
	FullName string
}

func NewAuthService(userRepo *repository.UserRepository, profileRepo *repository.ProfileRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, profileRepo: profileRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput) (string, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return "", apperror.ErrPasswordHashingFailed
	}

	user, err := s.userRepo.CreateUser(ctx, input.Email, hashedPassword)
	if err != nil {
		if dbutils.IsUniqueViolation(err) {
			return "", apperror.ErrEmailAlreadyExists
		}
		return "", apperror.ErrInternal
	}

	_, err = s.profileRepo.CreateProfile(ctx, repository.CreateProfileInput{UserID: user.ID, FullName: input.FullName})
	if err != nil {
		return "", apperror.ErrProfileCreationFailed
	}

	return utils.GenerateToken(user.ID.String(), s.jwtSecret)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", apperror.ErrUserWithEmailNotFound
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", apperror.ErrInvalidCredentials
	}

	return utils.GenerateToken(user.ID.String(), s.jwtSecret)
}
