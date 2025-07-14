package services

import (
	"context"
	"errors"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UsersService struct {
	repo interfaces.UsersRepository
}

func NewUsersService(repo interfaces.UsersRepository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) Register(ctx context.Context, cond models.User) error {
	err := s.repo.Register(ctx, cond)
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersService) Login(ctx context.Context, Login string) (*models.User, error) {
	result, err := s.repo.Login(ctx, Login)
	if err != nil {
		return nil, err
	}
	return result, nil
}
