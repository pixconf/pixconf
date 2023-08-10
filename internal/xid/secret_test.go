package xid

import (
	"testing"
)

func TestGenerateSecretID(t *testing.T) {
	char, err := GenerateSecretID()
	if err != nil && char != "" {
		t.Error("secret id must be empty on error")
	}

	if err != nil {
		t.Error(err)
	}

	if len(char) != SecretIDSize {
		t.Error("wrong len of secret id")
	}
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
