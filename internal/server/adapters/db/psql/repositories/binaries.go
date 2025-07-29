package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/models"
	"main/internal/server/services"
)

// BinariesRepository implements the binaries data access layer for PostgreSQL
type BinariesRepository struct {
	db *psql.DB
}

// NewBinariesRepository creates a new BinariesRepository instance
func NewBinariesRepository(db *psql.DB) *BinariesRepository {
	return &BinariesRepository{
		db: db,
	}
}

// Get retrieves binary data by title and user ID from the database
func (r *BinariesRepository) Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error) {
	var result models.BinaryData

	err := r.db.Conn.QueryRowContext(ctx, stmt.binary.get, title, UserID).Scan(&result.ID, &result.Title, &result.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrBinaryNotFound
		}
		return nil, err
	}
	return &result, nil
}

// Add stores new binary data in the database
func (r *BinariesRepository) Add(ctx context.Context, cond models.BinaryData) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.binary.add, cond.Title, cond.UserID, cond.Data).Scan(&title)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return "", services.ErrBinaryAlreadyExists
	}

	if err != nil {
		return "", err
	}
	return title, nil
}

// Update modifies existing binary data in the database
func (r *BinariesRepository) Update(ctx context.Context, cond models.BinaryData) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.binary.update, cond.Data, cond.Title, cond.UserID).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

// Delete removes binary data from the database by title and user ID
func (r *BinariesRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.binary.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
