// Package crypto provides cryptographic utilities for the server application.
// It includes symmetric encryption using AES-GCM and password hashing using bcrypt.
//
// The package contains the following main components:
//
//   - Aes: Structure for AES-GCM encryption and decryption.
//     Methods:
//
//   - Encrypt(data []byte) ([]byte, error): Encrypts data and prepends the nonce.
//
//   - Decrypt(ciphertext []byte) ([]byte, error): Extracts the nonce and decrypts data.
//
//   - PassCrypto: Structure for password hashing and verification.
//     Methods:
//
//   - HashPassword(password string) (string, error): Hashes a plain text password using bcrypt.
//
//   - CheckPassword(password string, hashedPassword string) error: Compares a plain text password with its hashed version.
//
// The configuration for the crypto package is provided via a secret key for AES operations.
package crypto
