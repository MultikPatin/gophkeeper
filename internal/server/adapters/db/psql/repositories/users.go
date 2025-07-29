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

// UsersRepository handles all database operations related to user management.
type UsersRepository struct {
	db *psql.DB // Database connection instance used for executing queries
}

// NewUsersRepository creates a new instance of UsersRepository with provided DB connection.
func NewUsersRepository(db *psql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

// Register registers a new user into the system using the given User model.
// It returns an error if registration fails or the login already exists.
func (r *UsersRepository) Register(ctx context.Context, cond models.User) (int64, error) {
	var userID int64

	err := r.db.Conn.QueryRowContext(ctx, stmt.user.register, cond.Login, cond.Password).Scan(&userID)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return -1, services.ErrLoginAlreadyExists
	}

	if err != nil {
		return -1, err
	}
	return userID, nil
}

// Login retrieves a user from the database based on their login.
// If no matching user is found, it returns an appropriate error message.
func (r *UsersRepository) Login(ctx context.Context, Login string) (*models.User, error) {
	var user models.User

	err := r.db.Conn.QueryRowContext(ctx, stmt.user.login, Login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, services.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
