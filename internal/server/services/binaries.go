package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

// BinariesService manages business logic for binary data storage and retrieval.
// It integrates with a repository for persistence and a crypto service for encryption/decryption.
type BinariesService struct {
	r interfaces.BinariesRepository // Repository for accessing binary data storage.
	c interfaces.CryptoService      // Service responsible for encryption and decryption.
}

// NewBinariesService instantiates a new BinariesService instance with dependencies injected.
func NewBinariesService(r interfaces.BinariesRepository, c interfaces.CryptoService) *BinariesService {
	return &BinariesService{
		r: r,
		c: c,
	}
}

// Get fetches a binary data item by title and user ID, then decrypts the content.
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

// Add inserts a new binary data item into storage after encrypting its contents.
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

// Update modifies an existing binary data item, first encrypting the updated content.
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

// Delete removes a binary data item identified by title and user ID.
func (s *BinariesService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}

// decrypt takes a binary data item and decrypts its content using the configured crypto service.
func (s *BinariesService) decrypt(result *models.BinaryData) (*models.BinaryData, error) {
	var err error

	result.Data, err = s.c.Decrypt(result.Data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// encrypt encrypts the binary data content prior to persisting it.
func (s *BinariesService) encrypt(cond models.BinaryData) (models.BinaryData, error) {
	var err error

	cond.Data, err = s.c.Encrypt(cond.Data)
	if err != nil {
		return models.BinaryData{}, err
	}

	return cond, nil
}
