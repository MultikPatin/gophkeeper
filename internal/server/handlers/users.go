package handlers

import (
	"main/internal/server/interfaces"
)

type UsersHandler struct {
	s interfaces.BinariesService
}

func NewUsersHandler(s interfaces.BinariesService) *UsersHandler {
	return &UsersHandler{
		s: s,
	}
}
