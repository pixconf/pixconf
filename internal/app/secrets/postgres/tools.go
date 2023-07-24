package postgres

import (
	"crypto/rand"
	"math/big"
)

const (
	SecretIDSize   = 24
	SecretIDSChars = "abcdefghjkmnpqrstuvwxyz123456789"
)

func GenerateSecretID() (string, error) {
	randomString := make([]byte, SecretIDSize)

	charSetLength := big.NewInt(int64(len(SecretIDSChars)))
	for i := 0; i < SecretIDSize; i++ {
		randomIndex, err := rand.Int(rand.Reader, charSetLength)
		if err != nil {
			return "", err
		}
		randomString[i] = SecretIDSChars[randomIndex.Int64()]
	}

	return string(randomString), nil
}
