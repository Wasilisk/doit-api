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
		return "", apperror.New(apperror.CodePasswordHashingFailed)
	}

	user, err := s.userRepo.CreateUser(ctx, input.Email, hashedPassword)
	if err != nil {
		if dbutils.IsUniqueViolation(err) {
			return "", apperror.New(apperror.CodeEmailAlreadyExists)
		}
		return "", apperror.New(apperror.CodeInternal)
	}

	_, err = s.profileRepo.CreateProfile(ctx, repository.CreateProfileInput{UserID: user.ID, FullName: input.FullName})
	if err != nil {
		return "", apperror.New(apperror.CodeProfileCreationFailed)
	}

	return utils.GenerateToken(user.ID.String(), s.jwtSecret)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", apperror.New(apperror.CodeUserWithEmailNotFound)
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", apperror.New(apperror.CodeInvalidCredentials)
	}

	return utils.GenerateToken(user.ID.String(), s.jwtSecret)
}
