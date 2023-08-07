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

func TestIsValidSecretIDKey(t *testing.T) {
	keys := map[string]bool{
		"ab12efghjkmn56789pqrstuv":     true,
		"abc123def456ghj789kmnpqr":     true,
		"ab12efghjkmn56789pqrstuvw":    false,
		"abcdefghjkmnpqrstuvwxyz":      false,
		"AB12EFGHJKMN56789PQRSTUV":     false,
		"abc123def456ghj789kmnpqrs!":   false,
		"abcdefghjkmnpqrstuvwxyz12345": false,
	}

	for key, resp := range keys {
		if ok := IsValidSecretIDKey(key); ok != resp {
			t.Errorf("wrong validate: got %#v, valid %#v, key %s", ok, resp, key)
		}
	}
}
