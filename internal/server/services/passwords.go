package services

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
)

type PasswordsService struct {
	r interfaces.PasswordsRepository
	c interfaces.CryptoService
}

func NewPasswordsService(r interfaces.PasswordsRepository, c interfaces.CryptoService) *PasswordsService {
	return &PasswordsService{
		r: r,
		c: c,
	}
}

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

func (s *PasswordsService) Delete(ctx context.Context, title string, UserID int64) error {
	err := s.r.Delete(ctx, title, UserID)
	if err != nil {
		return err
	}
	return nil
}

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

func (s *PasswordsService) encrypt(cond models.Password) (models.Password, error) {
	var err error

	cond.Login, err = s.c.Decrypt(cond.Login)
	if err != nil {
		return models.Password{}, err
	}
	cond.Password, err = s.c.Decrypt(cond.Password)
	if err != nil {
		return models.Password{}, err
	}

	return cond, nil
}
