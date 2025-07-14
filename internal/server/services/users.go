package services

import (
	"errors"
	"main/internal/server/interfaces"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UsersService struct {
	db interfaces.UsersRepository
}

func NewUsersService(db interfaces.UsersRepository) *UsersService {
	return &UsersService{
		db: db,
	}
}
