package interfaces

import (
	"context"
	"main/internal/server/models"
)

type BinariesRepository interface {
	Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error)
	Add(ctx context.Context, cond models.BinaryData) (string, error)
	Update(ctx context.Context, cond models.BinaryData) (string, error)
	Delete(ctx context.Context, title string, UserID int64) error
}

type PasswordsRepository interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Password, error)
	Add(ctx context.Context, cond models.Password) (string, error)
	Update(ctx context.Context, cond models.Password) (string, error)
	Delete(ctx context.Context, title string, UserID int64) error
}

type CardsRepository interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Card, error)
	Add(ctx context.Context, cond models.Card) (string, error)
	Update(ctx context.Context, cond models.Card) (string, error)
	Delete(ctx context.Context, title string, UserID int64) error
}

type UsersRepository interface {
	Register(ctx context.Context, cond models.User) error
	Login(ctx context.Context, Login string) (*models.User, error)
}
