package xid

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (er errorReader) Read(p []byte) (n int, err error) {
	return 0, assert.AnError
}

func TestGenerateSecretID(t *testing.T) {
	secretID, err := GenerateSecretID()
	if err != nil && secretID != "" {
		t.Error("secret id must be empty on error")
	}

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, SecretIDSize, len(secretID), "Generated SecretID length should match SecretIDSize")

	for _, char := range secretID {
		assert.Contains(t, SecretIDSChars, string(char), "Generated SecretID contains invalid character")
	}
}

func TestGenerateSecretID_Error(t *testing.T) {
	rand.Reader = errorReader{}

	secretID, err := GenerateSecretID()
	if err == nil {
		t.Errorf("GenerateSecretID should have returned an error")
	}

	assert.Empty(t, secretID, "Generated SecretID should be empty on error")
}

func TestIsValidSecretID(t *testing.T) {
	keys := map[string]bool{
		"sec-ab12efghjkmn56789pqrstuv": true,
		"sec-abc123def456ghj789kmnpqr": true,
		"abc123def456ghj789kmnpqr":     false,
		"ab12efghjkmn56789pqrstuvw":    false,
		"sec-abcdefghjkmnpqrstuvwxyz":  false,
		"AB12EFGHJKMN56789PQRSTUV":     false,
		"abc123def456ghj789kmnpqrs!":   false,
		"abcdefghjkmnpqrstuvwxyz12345": false,
	}

	for key, resp := range keys {
		if ok := IsValidSecretID(key); ok != resp {
			t.Errorf("wrong validate: got %#v, valid %#v, key %s", ok, resp, key)
		}
	}
}

func TestGetPublicSecretID(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"ab12efghjkmn56789pqrstuv", "sec-ab12efghjkmn56789pqrstuv"},
		{"test123", "sec-test123"},
	}

	for _, testCase := range testCases {
		output := GetPublicSecretID(testCase.input)
		if output != testCase.expectedOutput {
			t.Errorf("For input %s, expected %s, but got %s", testCase.input, testCase.expectedOutput, output)
		}
	}
}

func TestGetPrivateSecretID(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"sec-ab12efghjkmn56789pqrstuv", "ab12efghjkmn56789pqrstuv"},
		{"sec-test123", "test123"},
		{"no_prefix", "no_prefix"},
	}

	for _, testCase := range testCases {
		output := GetPrivateSecretID(testCase.input)
		if output != testCase.expectedOutput {
			t.Errorf("For input %s, expected %s, but got %s", testCase.input, testCase.expectedOutput, output)
		}
	}
}
