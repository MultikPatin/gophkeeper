package services

import (
	"main/internal/server/interfaces"
)

type PasswordsService struct {
	db interfaces.BinariesRepository
}

func NewPasswordsService(db interfaces.BinariesRepository) *PasswordsService {
	return &PasswordsService{
		db: db,
	}
}
