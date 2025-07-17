package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type PassCrypto struct {
}

func NewPassCrypto() (*PassCrypto, error) {
	return &PassCrypto{}, nil
}

func (s *PassCrypto) IsEqual(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *PassCrypto) Hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
