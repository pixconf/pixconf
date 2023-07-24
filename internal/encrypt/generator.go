package encrypt

import (
	"crypto/rand"
	"io"
)

func GenerateKey() ([]byte, error) {
	key := make([]byte, KeySize)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}

	return key, nil
}
