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
