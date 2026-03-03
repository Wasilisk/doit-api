package service

import (
	"context"
	"errors"

	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/utils"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(ctx context.Context, email, password string) (string, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.CreateUser(ctx, email, hashedPassword)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return utils.GenerateToken(user.ID, s.jwtSecret)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user with this email does not exist")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID, s.jwtSecret)
}
