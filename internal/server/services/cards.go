package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

type CardsService struct {
	r interfaces.CardsRepository
	c interfaces.CryptoService
}

func NewCardsService(r interfaces.CardsRepository, c interfaces.CryptoService) *CardsService {
	return &CardsService{
		r: r,
		c: c,
	}
}

func (s *CardsService) Get(ctx context.Context, title string, UserID int64) (*models.Card, error) {
	result, err := s.r.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CardsService) Add(ctx context.Context, cond models.Card) (int64, error) {
	result, err := s.r.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *CardsService) Update(ctx context.Context, cond models.Card) (int64, error) {
	result, err := s.r.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *CardsService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
