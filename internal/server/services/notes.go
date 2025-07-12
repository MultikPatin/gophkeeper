package services

import (
	"main/internal/server/interfaces"
)

type NotesService struct {
	db *interfaces.NotesRepository
}

func NewNotesService(db *interfaces.NotesRepository) *NotesService {
	return &NotesService{
		db: db,
	}
}
