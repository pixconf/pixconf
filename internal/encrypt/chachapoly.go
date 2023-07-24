package encrypt

import (
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

type ChachaPoly struct {
	aead cipher.AEAD
}

func NewChachaPoly(key []byte) (Encrypter, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	return &ChachaPoly{aead: aead}, nil
}

func (e *ChachaPoly) GetEncryptType() Type {
	return TypeChachaPoly
}

func (e *ChachaPoly) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, chacha20poly1305.NonceSizeX)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return e.aead.Seal(nonce, nonce, plaintext, nil), nil
}

func (e *ChachaPoly) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < chacha20poly1305.NonceSizeX {
		return nil, ErrMalformed
	}

	nonce := ciphertext[:chacha20poly1305.NonceSizeX]
	cleanCiphertext := ciphertext[chacha20poly1305.NonceSizeX:]

	return e.aead.Open(nil, nonce, cleanCiphertext, nil)
}
