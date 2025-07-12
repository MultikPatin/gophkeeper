package services

import (
	"main/internal/server/interfaces"
)

type BinariesService struct {
	db *interfaces.BinariesRepository
}

func NewBinariesService(db *interfaces.BinariesRepository) *BinariesService {
	return &BinariesService{
		db: db,
	}
}
