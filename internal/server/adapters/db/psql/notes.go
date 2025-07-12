package psql

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type NotesRepository struct {
	db *DB
}

func NewNotesRepository(db *DB) *NotesRepository {
	return &NotesRepository{
		db: db,
	}
}
