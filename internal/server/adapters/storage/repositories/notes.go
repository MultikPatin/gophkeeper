package repositories

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/interfaces"
)

type NotesRepository struct {
	db *interfaces.DB
}

func NewNotesRepository(db *interfaces.DB) *NotesRepository {
	return &NotesRepository{
		db: db,
	}
}
