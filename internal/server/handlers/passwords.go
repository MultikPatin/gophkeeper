package handlers

import (
	"main/internal/server/interfaces"
)

type PasswordsHandler struct {
	s interfaces.BinariesService
}

func NewPasswordsHandler(s interfaces.BinariesService) *PasswordsHandler {
	return &PasswordsHandler{
		s: s,
	}
}
