package services

import (
	"context"
	"errors"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	"time"
)

// Error definitions for common scenarios in user service operations.
var (
	ErrLoginAlreadyExists = errors.New("login already exists")           // Thrown when attempting to register a duplicate login.
	ErrInvalidCredentials = errors.New("login or password is not valid") // Raised when invalid credentials are presented during login.
	ErrUserNotFound       = errors.New("user not found")                 // Raised when attempting to authenticate a non-existent user.
)

// UsersService encapsulates user-related business logic, handling registration and authentication processes.
type UsersService struct {
	r interfaces.UsersRepository   // Dependency for interacting with the user repository.
	c interfaces.PassCryptoService // Dependency for password hashing and verification.
}

// NewUsersService creates a new instance of UsersService with the necessary dependencies.
func NewUsersService(r interfaces.UsersRepository, c interfaces.PassCryptoService) *UsersService {
	return &UsersService{
		r: r,
		c: c,
	}
}

// Register performs user registration, hashing the provided password and persisting the user data.
func (s *UsersService) Register(ctx context.Context, cond models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // Set timeout for the operation.
	defer cancel()

	hash, err := s.c.Hash(cond.Password)
	if err != nil {
		return -1, err
	}
	cond.Password = hash

	userID, err := s.r.Register(ctx, cond)
	if err != nil {
		return -1, err
	}
	return userID, nil
}

// Login authenticates a user by validating their credentials against persisted data.
func (s *UsersService) Login(ctx context.Context, cond models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // Set timeout for the operation.
	defer cancel()

	result, err := s.r.Login(ctx, cond.Login)
	if err != nil {
		return -1, err
	}

	err = s.c.IsEqual(cond.Password, result.Password)
	if err != nil {
		return -1, ErrInvalidCredentials
	}

	return result.ID, nil
}
