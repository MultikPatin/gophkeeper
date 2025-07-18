package interfaces

// CryptoService defines the interface for encrypting and decrypting arbitrary byte slices.
// Implementations must support two-way encryption for secure communication/storage of sensitive data.
type CryptoService interface {
	Encrypt([]byte) ([]byte, error) // Encrypts plain text data returning the encrypted result along with any potential errors.
	Decrypt([]byte) ([]byte, error) // Decrypts encrypted data back to its original form.
}

// PassCryptoService outlines the contract for handling password-related security operations.
// It includes methods for comparing passwords against hashed versions and creating new hashes.
type PassCryptoService interface {
	IsEqual(password string, hash string) error // Verifies if a clear-text password matches its hashed counterpart.
	Hash(password string) (string, error)       // Generates a secure hash for a given password.
}

// JWTService defines methods for validating and generating JSON Web Tokens (JWT).
// Used primarily for authenticating requests between client-server communications.
type JWTService interface {
	Verify(tokenStr string) (int64, error) // Validates a JWT token extracting the associated user ID.
	Generate(userID int64) (string, error) // Creates a new JWT token tied to a specific user ID.
}
