package repositories

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/interfaces"
)

type CardsRepository struct {
	db *interfaces.DB
}

func NewCardsRepository(db *interfaces.DB) *CardsRepository {
	return &CardsRepository{
		db: db,
	}
}
