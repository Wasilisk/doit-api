package service

import (
	"context"
	"errors"

	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/utils"
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
		return "", err
	}

	user, err := s.userRepo.CreateUser(ctx, input.Email, hashedPassword)
	if err != nil {
		return "", errors.New(err.Error())
	}

	_, err = s.profileRepo.CreateProfile(ctx, repository.CreateProfileInput{UserID: user.ID, FullName: input.FullName})
	if err != nil {
		return "", errors.New("failed to create profile")
	}

	return utils.GenerateToken(user.ID.String(), s.jwtSecret)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user with this email does not exist")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID.String(), s.jwtSecret)
}
