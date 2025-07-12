package repositories

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/interfaces"
)

type BinariesRepository struct {
	db *interfaces.DB
}

func NewBinariesRepository(db *interfaces.DB) *BinariesRepository {
	return &BinariesRepository{
		db: db,
	}
}
