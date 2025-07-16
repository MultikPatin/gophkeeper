package services

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	"time"
)

var (
	ErrLoginAlreadyExists        = errors.New("login already exists")
	ErrAuthCredentialsIsNotValid = errors.New("login or password is not valid")
)

type UsersService struct {
	repo interfaces.UsersRepository
}

func NewUsersService(repo interfaces.UsersRepository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) Register(ctx context.Context, cond models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	hash, err := hashPassword(cond.Password)
	if err != nil {
		return err
	}
	cond.Password = hash

	err = s.repo.Register(ctx, cond)
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersService) Login(ctx context.Context, cond models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := s.repo.Login(ctx, cond.Login)
	if err != nil {
		return err
	}

	if !isEqualPasswords(cond.Password, result.Password) {
		return ErrAuthCredentialsIsNotValid
	}

	return nil
}

func isEqualPasswords(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
