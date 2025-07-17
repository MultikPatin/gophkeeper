package interfaces

type CryptoService interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type PassCryptoService interface {
	IsEqual(password string, hash string) bool
	Hash(password string) (string, error)
}

type JWTService interface {
	Verify(tokenStr string) (int64, error)
	Generate(userID int64) (string, error)
}
