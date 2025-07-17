package services

import (
	"context"
	"errors"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	"time"
)

var (
	ErrLoginAlreadyExists        = errors.New("login already exists")
	ErrAuthCredentialsIsNotValid = errors.New("login or password is not valid")
)

type UsersService struct {
	r interfaces.UsersRepository
	c interfaces.PassCryptoService
}

func NewUsersService(r interfaces.UsersRepository, c interfaces.PassCryptoService) *UsersService {
	return &UsersService{
		r: r,
		c: c,
	}
}

func (s *UsersService) Register(ctx context.Context, cond models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	hash, err := s.c.Hash(cond.Password)
	if err != nil {
		return err
	}
	cond.Password = hash

	err = s.r.Register(ctx, cond)
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersService) Login(ctx context.Context, cond models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := s.r.Login(ctx, cond.Login)
	if err != nil {
		return err
	}

	if !s.c.IsEqual(cond.Password, result.Password) {
		return ErrAuthCredentialsIsNotValid
	}

	return nil
}
