package psql

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/models"
)

type PasswordsRepository struct {
	db *DB
}

func NewPasswordsRepository(db *DB) *PasswordsRepository {
	return &PasswordsRepository{
		db: db,
	}
}

func (r *PasswordsRepository) Get(ctx context.Context, title string, UserID int64) (*models.Password, error) {
	var result models.Password

	err := r.db.conn.QueryRowContext(ctx, stmt.password.get, title, UserID).Scan(&result.ID, &result.Title, &result.Login, &result.Password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *PasswordsRepository) Add(ctx context.Context, cond models.Password) (int64, error) {
	var ID int64

	err := r.db.conn.QueryRowContext(ctx, stmt.password.add, cond.Title, cond.UserID, cond.Login, cond.Password).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *PasswordsRepository) Update(ctx context.Context, cond models.Password) (int64, error) {
	var ID int64

	err := r.db.conn.QueryRowContext(ctx, stmt.password.update, cond.Login, cond.Password, cond.Title, cond.UserID).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *PasswordsRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.conn.ExecContext(ctx, stmt.password.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
