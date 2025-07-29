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

// CardsRepository implements the cards data access layer for PostgreSQL
type CardsRepository struct {
	db *psql.DB
}

// NewCardsRepository creates a new CardsRepository instance
func NewCardsRepository(db *psql.DB) *CardsRepository {
	return &CardsRepository{
		db: db,
	}
}

// Get retrieves credit card information by title and user ID from the database
func (r *CardsRepository) Get(ctx context.Context, title string, UserID int64) (*models.Card, error) {
	var result models.Card

	err := r.db.Conn.QueryRowContext(ctx, stmt.card.get, title, UserID).Scan(&result.ID, &result.Title, &result.Bank, &result.Number, &result.DataEnd, &result.SecretCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrCardNotFound
		}
		return nil, err
	}
	return &result, nil
}

// Add stores new credit card information in the database
func (r *CardsRepository) Add(ctx context.Context, cond models.Card) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.card.add, cond.Title, cond.UserID, cond.Bank, cond.Number, cond.DataEnd, cond.SecretCode).Scan(&title)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return "", services.ErrCardAlreadyExists
	}

	if err != nil {
		return "", err
	}
	return title, nil
}

// Update modifies existing credit card information in the database
func (r *CardsRepository) Update(ctx context.Context, cond models.Card) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.card.update, cond.Bank, cond.Number, cond.DataEnd, cond.SecretCode, cond.Title, cond.UserID).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

// Delete removes credit card information from the database by title and user ID
func (r *CardsRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.card.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
