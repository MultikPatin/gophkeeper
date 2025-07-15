package psql

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/models"
)

type BinariesRepository struct {
	db *DB
}

func NewBinariesRepository(db *DB) *BinariesRepository {
	return &BinariesRepository{
		db: db,
	}
}

func (r *BinariesRepository) Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error) {
	var result models.BinaryData

	err := r.db.conn.QueryRowContext(ctx, stmt.binary.get, title, UserID).Scan(&result.ID, &result.Title, &result.Data)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BinariesRepository) Add(ctx context.Context, cond models.BinaryData) (int64, error) {
	var ID int64

	err := r.db.conn.QueryRowContext(ctx, stmt.binary.add, cond.Title, cond.UserID, cond.Data).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *BinariesRepository) Update(ctx context.Context, cond models.BinaryData) (int64, error) {
	var ID int64

	err := r.db.conn.QueryRowContext(ctx, stmt.binary.update, cond.Data).Scan(&ID)
	if err != nil {
		return -1, err
	}
	return ID, nil
}

func (r *BinariesRepository) Delete(ctx context.Context, ID int64) error {
	_, err := r.db.conn.ExecContext(ctx, stmt.binary.delete, ID)
	if err != nil {
		return err
	}
	return nil
}
