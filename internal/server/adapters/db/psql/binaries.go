package psql

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type BinariesRepository struct {
	db *DB
}

func NewBinariesRepository(db *DB) *BinariesRepository {
	return &BinariesRepository{
		db: db,
	}
}
