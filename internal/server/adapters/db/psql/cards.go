package psql

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type CardsRepository struct {
	db *DB
}

func NewCardsRepository(db *DB) *CardsRepository {
	return &CardsRepository{
		db: db,
	}
}
