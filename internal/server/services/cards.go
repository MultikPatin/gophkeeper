package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

// CardsService manages the lifecycle of credit card entities, integrating encryption for sensitive data.
type CardsService struct {
	r interfaces.CardsRepository // Repository dependency for interacting with the persistent store.
	c interfaces.CryptoService   // Encryption service dependency for securing card data.
}

// NewCardsService creates a new instance of CardsService with injected dependencies.
func NewCardsService(r interfaces.CardsRepository, c interfaces.CryptoService) *CardsService {
	return &CardsService{
		r: r,
		c: c,
	}
}

// Get retrieves a credit card by title and user ID, decrypting its confidential fields.
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

// Add persists a new credit card, encrypting its sensitive fields beforehand.
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

// Update updates an existing credit card record, re-encrypting modified fields.
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

// Delete deletes a credit card record by title and user ID without needing decryption.
func (s *CardsService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}

// decrypt deobfuscates encrypted fields of a credit card entity.
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

// encrypt secures the sensitive fields of a credit card entity before storage.
func (s *CardsService) encrypt(cond models.Card) (models.Card, error) {
	var err error

	cond.Bank, err = s.c.Encrypt(cond.Bank)
	if err != nil {
		return models.Card{}, err
	}
	cond.Number, err = s.c.Encrypt(cond.Number)
	if err != nil {
		return models.Card{}, err
	}
	cond.DataEnd, err = s.c.Encrypt(cond.DataEnd)
	if err != nil {
		return models.Card{}, err
	}
	cond.SecretCode, err = s.c.Encrypt(cond.SecretCode)
	if err != nil {
		return models.Card{}, err
	}

	return cond, nil
}
