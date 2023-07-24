package encrypt

import (
	"encoding/base64"
	"errors"
)

const KeySize = 32

var (
	ErrKeySize            = errors.New("encryption key must be 32 bytes")
	ErrMalformed          = errors.New("malformed ciphertext")
	ErrUnknownEncryptType = errors.New("unknown encrypt type")
)

type Type uint8

const (
	TypeAesGCM     = 0
	TypeChachaPoly = 1
)

type Encrypter interface {
	GetEncryptType() Type
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

func NewEncoded(key string, encryptType Type) (Encrypter, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return New(decodedKey, encryptType)
}

func New(key []byte, encryptType Type) (Encrypter, error) {
	if len(key) != KeySize {
		return nil, ErrKeySize
	}

	switch encryptType {
	case TypeAesGCM:
		return NewAesGCM(key)

	case TypeChachaPoly:
		return NewChachaPoly(key)

	default:
		return nil, ErrUnknownEncryptType
	}
}
