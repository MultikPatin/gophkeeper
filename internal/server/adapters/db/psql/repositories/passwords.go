package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/models"
)

// PasswordsRepository implements the passwords data access layer for PostgreSQL
type PasswordsRepository struct {
	db *psql.DB // Database connection
}

// NewPasswordsRepository creates a new PasswordsRepository instance
func NewPasswordsRepository(db *psql.DB) *PasswordsRepository {
	return &PasswordsRepository{
		db: db,
	}
}

// Get retrieves password information by title and user ID from the database
func (r *PasswordsRepository) Get(ctx context.Context, title string, UserID int64) (*models.Password, error) {
	var result models.Password

	err := r.db.Conn.QueryRowContext(ctx, stmt.password.get, title, UserID).Scan(&result.ID, &result.Title, &result.Login, &result.Password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Add stores new password information in the database
func (r *PasswordsRepository) Add(ctx context.Context, cond models.Password) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.password.add, cond.Title, cond.UserID, cond.Login, cond.Password).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

// Update modifies existing password information in the database
func (r *PasswordsRepository) Update(ctx context.Context, cond models.Password) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.password.update, cond.Login, cond.Password, cond.Title, cond.UserID).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

// Delete removes password information from the database by title and user ID
func (r *PasswordsRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.password.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
