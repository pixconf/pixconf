package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AesGCM struct {
	block cipher.Block
}

func NewAesGCM(key []byte) (Encrypter, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AesGCM{block: block}, nil
}

func (e *AesGCM) GetEncryptType() Type {
	return TypeAesGCM
}

func (e *AesGCM) Encrypt(data []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func (e *AesGCM) Decrypt(data []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}

	if len(data) < gcm.NonceSize() {
		return nil, ErrMalformed
	}

	nonce := data[:gcm.NonceSize()]
	cleanCiphertext := data[gcm.NonceSize():]

	return gcm.Open(nil, nonce, cleanCiphertext, nil)
}
