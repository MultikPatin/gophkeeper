package repositories

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/interfaces"
)

type UsersRepository struct {
	db *interfaces.DB
}

func NewUsersRepository(db *interfaces.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}
