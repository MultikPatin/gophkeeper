package handlers

import (
	"main/internal/server/interfaces"
)

type CardsHandler struct {
	s interfaces.BinariesService
}

func NewCardsHandler(s interfaces.BinariesService) *CardsHandler {
	return &CardsHandler{
		s: s,
	}
}
