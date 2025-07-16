package handlers

import (
	"main/internal/server/interfaces"
)

type CardsHandler struct {
	s interfaces.CardsService
}

func NewCardsHandler(s interfaces.CardsService) *CardsHandler {
	return &CardsHandler{
		s: s,
	}
}
