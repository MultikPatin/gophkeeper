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

	result, err = s.decrypt(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CardsService) Add(ctx context.Context, cond models.Card) (string, error) {
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

func (s *CardsService) Update(ctx context.Context, cond models.Card) (string, error) {
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

func (s *CardsService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CardsService) decrypt(result *models.Card) (*models.Card, error) {
	var err error

	result.Bank, err = s.c.Decrypt(result.Bank)
	if err != nil {
		return nil, err
	}
	result.Number, err = s.c.Decrypt(result.Number)
	if err != nil {
		return nil, err
	}
	result.DataEnd, err = s.c.Decrypt(result.DataEnd)
	if err != nil {
		return nil, err
	}
	result.SecretCode, err = s.c.Decrypt(result.SecretCode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CardsService) encrypt(cond models.Card) (models.Card, error) {
	var err error

	cond.Bank, err = s.c.Decrypt(cond.Bank)
	if err != nil {
		return models.Card{}, err
	}
	cond.Number, err = s.c.Decrypt(cond.Number)
	if err != nil {
		return models.Card{}, err
	}
	cond.DataEnd, err = s.c.Decrypt(cond.DataEnd)
	if err != nil {
		return models.Card{}, err
	}
	cond.SecretCode, err = s.c.Decrypt(cond.SecretCode)
	if err != nil {
		return models.Card{}, err
	}

	return cond, nil
}
