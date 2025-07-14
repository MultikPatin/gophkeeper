package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

type BinariesService struct {
	repo interfaces.BinariesRepository
}

func NewBinariesService(repo interfaces.BinariesRepository) *BinariesService {
	return &BinariesService{
		repo: repo,
	}
}

func (s *BinariesService) Get(ctx context.Context, title string, owner int64) (*models.BinaryData, error) {
	result, err := s.repo.Get(ctx, title, owner)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *BinariesService) Add(ctx context.Context, cond models.BinaryData) (int64, error) {
	result, err := s.repo.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *BinariesService) Update(ctx context.Context, cond models.BinaryData) (int64, error) {
	result, err := s.repo.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *BinariesService) Delete(ctx context.Context, ID int64) error {
	err := s.repo.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
