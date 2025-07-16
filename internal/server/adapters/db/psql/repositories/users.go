package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/models"
	"main/internal/server/services"
)

type UsersRepository struct {
	db *psql.DB
}

func NewUsersRepository(db *psql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) Register(ctx context.Context, cond models.User) error {
	_, err := r.db.Conn.ExecContext(ctx, stmt.user.register, cond.Login, cond.Password)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return services.ErrUserAlreadyExists
	}

	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepository) Login(ctx context.Context, Login string) (*models.User, error) {
	var user models.User

	err := r.db.Conn.QueryRowContext(ctx, stmt.user.login, Login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
