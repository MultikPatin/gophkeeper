package handlers

import (
	"main/internal/server/interfaces"
)

type PasswordsHandler struct {
	s interfaces.PasswordsService
}

func NewPasswordsHandler(s interfaces.PasswordsService) *PasswordsHandler {
	return &PasswordsHandler{
		s: s,
	}
}
