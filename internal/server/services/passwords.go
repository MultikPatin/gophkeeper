package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

type PasswordsService struct {
	repo interfaces.PasswordsRepository
}

func NewPasswordsService(repo interfaces.PasswordsRepository) *PasswordsService {
	return &PasswordsService{
		repo: repo,
	}
}

func (s *PasswordsService) Get(ctx context.Context, title string, UserID int64) (*models.Password, error) {
	result, err := s.repo.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PasswordsService) Add(ctx context.Context, cond models.Password) (int64, error) {
	result, err := s.repo.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *PasswordsService) Update(ctx context.Context, cond models.Password) (int64, error) {
	result, err := s.repo.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *PasswordsService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.repo.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
