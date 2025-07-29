package interfaces

import (
	"context"
	"main/internal/server/models"
)

// BinariesRepository defines the repository-level interface for binary data management.
// It supports retrieval, addition, updating, and deletion of binary records associated with users.
type BinariesRepository interface {
	Get(ctx context.Context, title string, UserID int64) (*models.BinaryData, error) // Retrieves binary data by title and user ID.
	Add(ctx context.Context, cond models.BinaryData) (string, error)                 // Adds new binary data.
	Update(ctx context.Context, cond models.BinaryData) (string, error)              // Updates existing binary data.
	Delete(ctx context.Context, title string, UserID int64) error                    // Deletes binary data by title and user ID.
}

// PasswordsRepository outlines the interface for password data management.
// Provides methods for retrieving, adding, modifying, and removing password entries linked to users.
type PasswordsRepository interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Password, error) // Fetches password by title and user ID.
	Add(ctx context.Context, cond models.Password) (string, error)                 // Adds a new password entry.
	Update(ctx context.Context, cond models.Password) (string, error)              // Modifies an existing password entry.
	Delete(ctx context.Context, title string, UserID int64) error                  // Removes a password entry by title and user ID.
}

// CardsRepository specifies the repository-level interface for credit card data management.
// Supports fetching, inserting, updating, and deleting card records connected to users.
type CardsRepository interface {
	Get(ctx context.Context, title string, UserID int64) (*models.Card, error) // Obtains a credit card by title and user ID.
	Add(ctx context.Context, cond models.Card) (string, error)                 // Adds a new credit card entry.
	Update(ctx context.Context, cond models.Card) (string, error)              // Edits an existing credit card entry.
	Delete(ctx context.Context, title string, UserID int64) error              // Eliminates a credit card by title and user ID.
}

// UsersRepository defines the interface for user account management.
// Includes methods for registering new users and logging them in.
type UsersRepository interface {
	Register(ctx context.Context, cond models.User) (int64, error) // Registers a new user account.
	Login(ctx context.Context, Login string) (*models.User, error) // Logs in a user by checking their credentials.
}
