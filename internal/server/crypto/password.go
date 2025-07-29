// Package crypto provides password hashing functionality for secure user authentication
package crypto

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// PassCrypto provides password hashing and verification functionality
type PassCrypto struct{}

// NewPassCrypto creates a new PassCrypto instance
func NewPassCrypto() *PassCrypto {
	return &PassCrypto{}
}

// Hash hashes a plain text password using bcrypt
// Returns the hashed password string
func (p *PassCrypto) Hash(password string) (string, error) {
	if password == "" {
		return "", errors.New("empty password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// IsEqual compares a plain text password with its hashed version
// Returns nil if they match, or an error if they don't
func (p *PassCrypto) IsEqual(password string, hash string) error {
	if password == "" || hash == "" {
		return errors.New("empty password or hashed password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("password check failed: %w", err)
	}
	return nil
}
