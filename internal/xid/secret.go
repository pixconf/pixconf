package xid

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
)

const (
	SecretIDSize   = 24
	SecretIDSChars = "abcdefghjkmnpqrstuvwxyz123456789"
)

var secretIDPattern = regexp.MustCompile(fmt.Sprintf("^[%s]{%d}$", SecretIDSChars, SecretIDSize))

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

func IsValidKey(input string) bool {
	return secretIDPattern.MatchString(input)
}
