package handlers

import (
	"main/internal/server/interfaces"
)

type BinariesHandler struct {
	s interfaces.BinariesService
}

func NewBinariesHandler(s interfaces.BinariesService) *BinariesHandler {
	return &BinariesHandler{
		s: s,
	}
}
