package interfaces

import (
	"context"
	"main/internal/server/models"
)

// BinariesService defines the business logic layer for managing binary data entities.
// Implements methods for retrieving, adding, updating, and deleting binary resources.
type BinariesService interface {
	Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error) // Retrieves binary data by title and user ID.
	Add(ctx context.Context, cond models.BinaryData) (string, error)                 // Adds new binary resource.
	Update(ctx context.Context, cond models.BinaryData) (string, error)              // Updates existing binary resource.
	Delete(ctx context.Context, title string, UserID int64) error                    // Deletes binary resource by title and user ID.
}

// PasswordsService outlines the service-layer interface for password data management.
// Contains methods for getting, adding, updating, and removing password records belonging to users.
type PasswordsService interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Password, error) // Fetches password by title and user ID.
	Add(ctx context.Context, cond models.Password) (string, error)                 // Adds a new password entry.
	Update(ctx context.Context, cond models.Password) (string, error)              // Modifies an existing password entry.
	Delete(ctx context.Context, title string, UserID int64) error                  // Removes a password entry by title and user ID.
}

// CardsService specifies the business logic for credit card data management.
// Offers methods for obtaining, saving, editing, and erasing credit card records linked to users.
type CardsService interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Card, error) // Gets a credit card by title and user ID.
	Add(ctx context.Context, cond models.Card) (string, error)                 // Adds a new credit card entry.
	Update(ctx context.Context, cond models.Card) (string, error)              // Updates an existing credit card entry.
	Delete(ctx context.Context, title string, UserID int64) error              // Deletes a credit card entry by title and user ID.
}

// UsersService defines the service-level interface for user account management.
// Methods include registering new users and processing log-in attempts.
type UsersService interface {
	Register(ctx context.Context, cond models.User) (int64, error) // Registers a new user account.
	Login(ctx context.Context, cond models.User) (int64, error)    // Handles user log-in process.
}
