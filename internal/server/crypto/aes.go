package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Aes implements AES-based symmetric encryption and decryption functionality.
type Aes struct {
	gcm cipher.AEAD // Galois Counter Mode (GCM) interface for performing authenticated encryption.
}

// NewAes initializes a new Aes instance with the provided encryption key.
// It sets up AES block cipher mode and constructs a GCM cipher.
func NewAes(key []byte) (*Aes, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &Aes{aesGCM}, nil
}

// Encrypt encrypts input plaintext data using AES-GCM encryption scheme.
// It generates a random nonce, seals the data, and appends the nonce at the beginning of the encrypted output.
func (a *Aes) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, a.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	res := a.gcm.Seal(nil, nonce, data, nil)
	res = append(nonce, res...)

	return res, nil
}

// Decrypt decrypts previously encrypted data back to its original form.
// It separates the nonce from the ciphertext and uses it for decryption via GCM.
func (a *Aes) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := a.gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	res, err := a.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
