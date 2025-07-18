package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

// PasswordsService manages the lifecycle of password entities, incorporating encryption for sensitive fields.
type PasswordsService struct {
	r interfaces.PasswordsRepository // Dependency for interacting with the underlying password repository.
	c interfaces.CryptoService       // Encryption service for protecting sensitive password data.
}

// NewPasswordsService creates a new instance of PasswordsService with injected dependencies.
func NewPasswordsService(r interfaces.PasswordsRepository, c interfaces.CryptoService) *PasswordsService {
	return &PasswordsService{
		r: r,
		c: c,
	}
}

// Get retrieves a password by title and user ID, decrypting its sensitive fields.
func (s *PasswordsService) Get(ctx context.Context, title string, UserID int64) (*models.Password, error) {
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

// Add saves a new password entry, first encrypting its sensitive fields.
func (s *PasswordsService) Add(ctx context.Context, cond models.Password) (string, error) {
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

// Update modifies an existing password record, re-encrypting its sensitive fields.
func (s *PasswordsService) Update(ctx context.Context, cond models.Password) (string, error) {
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

// Delete permanently removes a password entry by title and user ID.
func (s *PasswordsService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}

// decrypt deobfuscates encrypted fields of a password entity.
func (s *PasswordsService) decrypt(result *models.Password) (*models.Password, error) {
	var err error

	result.Login, err = s.c.Decrypt(result.Login)
	if err != nil {
		return nil, err
	}
	result.Password, err = s.c.Decrypt(result.Password)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// encrypt secures the sensitive fields of a password entity before storage.
func (s *PasswordsService) encrypt(cond models.Password) (models.Password, error) {
	var err error

	cond.Login, err = s.c.Encrypt(cond.Login)
	if err != nil {
		return models.Password{}, err
	}
	cond.Password, err = s.c.Encrypt(cond.Password)
	if err != nil {
		return models.Password{}, err
	}

	return cond, nil
}
