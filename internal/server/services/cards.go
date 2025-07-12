package services

import (
	"main/internal/server/interfaces"
)

type CardsService struct {
	db interfaces.CardsRepository
}

func NewCardsService(db interfaces.CardsRepository) *CardsService {
	return &CardsService{
		db: db,
	}
}
