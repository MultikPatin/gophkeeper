package psql

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/models"
)

type CardsRepository struct {
	db *DB
}

func NewCardsRepository(db *DB) *CardsRepository {
	return &CardsRepository{
		db: db,
	}
}

func (r *CardsRepository) Get(ctx context.Context, title string, UserID int64) (*models.Card, error) {
	var result models.Card

	err := r.db.conn.QueryRowContext(ctx, stmt.card.get, title, UserID).Scan(&result.ID, &result.Title, &result.Bank, &result.Number, &result.DataEnd, &result.SecretCode)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *CardsRepository) Add(ctx context.Context, cond models.Card) (int64, error) {
	var ID int64

	err := r.db.conn.QueryRowContext(ctx, stmt.card.add, cond.Title, cond.UserID, cond.Bank, cond.Number, cond.DataEnd, cond.SecretCode).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *CardsRepository) Update(ctx context.Context, cond models.Card) (int64, error) {
	var ID int64

	err := r.db.conn.QueryRowContext(ctx, stmt.card.update, cond.Bank, cond.Number, cond.DataEnd, cond.SecretCode).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *CardsRepository) Delete(ctx context.Context, ID int64) error {
	_, err := r.db.conn.ExecContext(ctx, stmt.card.delete, ID)
	if err != nil {
		return err
	}
	return nil
}
