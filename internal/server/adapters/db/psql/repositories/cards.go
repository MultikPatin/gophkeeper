package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/models"
)

type CardsRepository struct {
	db *psql.DB
}

func NewCardsRepository(db *psql.DB) *CardsRepository {
	return &CardsRepository{
		db: db,
	}
}

func (r *CardsRepository) Get(ctx context.Context, title string, UserID int64) (*models.Card, error) {
	var result models.Card

	err := r.db.Conn.QueryRowContext(ctx, stmt.card.get, title, UserID).Scan(&result.ID, &result.Title, &result.Bank, &result.Number, &result.DataEnd, &result.SecretCode)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *CardsRepository) Add(ctx context.Context, cond models.Card) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.card.add, cond.Title, cond.UserID, cond.Bank, cond.Number, cond.DataEnd, cond.SecretCode).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

func (r *CardsRepository) Update(ctx context.Context, cond models.Card) (string, error) {
	var title string

	err := r.db.Conn.QueryRowContext(ctx, stmt.card.update, cond.Bank, cond.Number, cond.DataEnd, cond.SecretCode, cond.Title, cond.UserID).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

func (r *CardsRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.card.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
