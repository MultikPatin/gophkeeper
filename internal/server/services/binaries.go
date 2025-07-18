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

	result, err = s.decrypt(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *BinariesService) Add(ctx context.Context, cond models.BinaryData) (string, error) {
	var err error

	cond, err = s.encrypt(cond)
	if err != nil {
		return "", err
	}

	result, err := s.r.Add(ctx, cond)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *BinariesService) Update(ctx context.Context, cond models.BinaryData) (string, error) {
	var err error

	cond, err = s.encrypt(cond)
	if err != nil {
		return "", err
	}

	result, err := s.r.Update(ctx, cond)
	if err != nil {
		return "", err
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

func (s *BinariesService) decrypt(result *models.BinaryData) (*models.BinaryData, error) {
	var err error

	result.Data, err = s.c.Decrypt(result.Data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *BinariesService) encrypt(cond models.BinaryData) (models.BinaryData, error) {
	var err error

	cond.Data, err = s.c.Encrypt(cond.Data)
	if err != nil {
		return models.BinaryData{}, err
	}

	return cond, nil
}
