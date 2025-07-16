package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type CryptoAes struct {
	gcm cipher.AEAD
}

func NewCryptoAes(key []byte) (*CryptoAes, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &CryptoAes{aesGCM}, nil
}

func (a *CryptoAes) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, a.gcm.NonceSize()) //12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	res := a.gcm.Seal(nil, nonce, data, nil)
	res = append(nonce, res...)

	return res, nil
}

func (a *CryptoAes) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := a.gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	res, err := a.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
