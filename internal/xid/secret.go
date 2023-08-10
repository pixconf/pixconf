package xid

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

const (
	SecretIDSize   = 24
	SecretIDSChars = "abcdefghjkmnpqrstuvwxyz123456789"
	SecretIDPrefix = "sec-"
)

var secretIDPattern = regexp.MustCompile(fmt.Sprintf("^%s[%s]{%d}$", SecretIDPrefix, SecretIDSChars, SecretIDSize))

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

func IsValidSecretID(input string) bool {
	return secretIDPattern.MatchString(input)
}

func GetPublicSecretID(input string) string {
	return fmt.Sprintf("%s%s", SecretIDPrefix, input)
}

func GetPrivateSecretID(input string) string {
	return strings.TrimPrefix(input, SecretIDPrefix)
}
