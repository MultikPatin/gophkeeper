package services

import (
	"main/internal/server/interfaces"
)

type UsersService struct {
	db interfaces.UsersRepository
}

func NewUsersService(db interfaces.UsersRepository) *UsersService {
	return &UsersService{
		db: db,
	}
}
