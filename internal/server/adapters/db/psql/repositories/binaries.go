package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/models"
)

type BinariesRepository struct {
	db *psql.DB
}

func NewBinariesRepository(db *psql.DB) *BinariesRepository {
	return &BinariesRepository{
		db: db,
	}
}

func (r *BinariesRepository) Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error) {
	var result models.BinaryData

	err := r.db.Conn.QueryRowContext(ctx, stmt.binary.get, title, UserID).Scan(&result.ID, &result.Title, &result.Data)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BinariesRepository) Add(ctx context.Context, cond models.BinaryData) (int64, error) {
	var ID int64

	err := r.db.Conn.QueryRowContext(ctx, stmt.binary.add, cond.Title, cond.UserID, cond.Data).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *BinariesRepository) Update(ctx context.Context, cond models.BinaryData) (int64, error) {
	var ID int64

	err := r.db.Conn.QueryRowContext(ctx, stmt.binary.update, cond.Data, cond.Title, cond.UserID).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *BinariesRepository) Delete(ctx context.Context, title string, UserID int64) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.binary.delete, title, UserID)
	if err != nil {
		return err
	}
	return nil
}
