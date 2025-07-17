package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

type BinariesService struct {
	r interfaces.BinariesRepository
	c interfaces.CryptoService
}

func NewBinariesService(r interfaces.BinariesRepository, c interfaces.CryptoService) *BinariesService {
	return &BinariesService{
		r: r,
		c: c,
	}
}

func (s *BinariesService) Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error) {
	result, err := s.r.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *BinariesService) Add(ctx context.Context, cond models.BinaryData) (int64, error) {
	result, err := s.r.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *BinariesService) Update(ctx context.Context, cond models.BinaryData) (int64, error) {
	result, err := s.r.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *BinariesService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
