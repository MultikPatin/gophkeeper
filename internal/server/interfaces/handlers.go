package interfaces

import (
	"context"
	"main/internal/server/models"
)

type BinariesHandlers interface {
	Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error)
	Add(ctx context.Context, cond models.BinaryData) (int64, error)
	Update(ctx context.Context, cond models.BinaryData) (int64, error)
	Delete(ctx context.Context, ID int64) error
}

type PasswordsHandlers interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Password, error)
	Add(ctx context.Context, cond models.Password) (int64, error)
	Update(ctx context.Context, cond models.Password) (int64, error)
	Delete(ctx context.Context, ID int64) error
}

type CardsHandlers interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Card, error)
	Add(ctx context.Context, cond models.Card) (int64, error)
	Update(ctx context.Context, cond models.Card) (int64, error)
	Delete(ctx context.Context, ID int64) error
}

type UsersHandlers interface {
	Register(ctx context.Context, cond models.User) error
	Login(ctx context.Context, Login string) (*models.User, error)
}
