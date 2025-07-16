package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/models"
)

type PasswordsRepository struct {
	db *psql.DB
}

func NewPasswordsRepository(db *psql.DB) *PasswordsRepository {
	return &PasswordsRepository{
		db: db,
	}
}

func (r *PasswordsRepository) Get(ctx context.Context, title string, UserID int64) (*models.Password, error) {
	var result models.Password

	err := r.db.Conn.QueryRowContext(ctx, stmt.password.get, title, UserID).Scan(&result.ID, &result.Title, &result.Login, &result.Password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *PasswordsRepository) Add(ctx context.Context, cond models.Password) (int64, error) {
	var ID int64

	err := r.db.Conn.QueryRowContext(ctx, stmt.password.add, cond.Title, cond.UserID, cond.Login, cond.Password).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *PasswordsRepository) Update(ctx context.Context, cond models.Password) (int64, error) {
	var ID int64

	err := r.db.Conn.QueryRowContext(ctx, stmt.password.update, cond.Login, cond.Password, cond.Title, cond.UserID).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *PasswordsRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.password.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
