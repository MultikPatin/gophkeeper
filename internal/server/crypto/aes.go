package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Aes struct {
	gcm cipher.AEAD
}

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

func (a *Aes) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, a.gcm.NonceSize()) //12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	res := a.gcm.Seal(nil, nonce, data, nil)
	res = append(nonce, res...)

	return res, nil
}

func (a *Aes) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := a.gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	res, err := a.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
