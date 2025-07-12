package psql

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type UsersRepository struct {
	db *DB
}

func NewUsersRepository(db *DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}
