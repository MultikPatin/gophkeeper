package handlers

import (
	"main/internal/server/interfaces"
)

type UsersHandler struct {
	s interfaces.UsersService
}

func NewUsersHandler(s interfaces.UsersService) *UsersHandler {
	return &UsersHandler{
		s: s,
	}
}
