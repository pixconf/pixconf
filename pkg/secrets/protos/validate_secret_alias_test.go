package protos

import (
	"fmt"
	"testing"
)

func TestIsValidAlias(t *testing.T) {
	validAliases := []string{
		"agent/abc123/xyz",
		"group/def456/abc/123",
		"system/ghi789/qqq",
		"user/jkl012/sss",
		"user/too/many/segments/here",
	}

	invalidAliases := []string{
		"invalidalias",
		"agent/",
		"agent/007",
		"test/test/test/123",
	}

	for _, alias := range validAliases {
		if !IsValidAlias(alias) {
			t.Errorf("Expected valid alias, but got invalid: %s", alias)
		}
	}

	for _, alias := range invalidAliases {
		if IsValidAlias(alias) {
			t.Errorf("Expected invalid alias, but got valid: %s", alias)
		}
	}
}

func TestIsValidAliases(t *testing.T) {
	validInput := map[string]SecretAlias{
		"agent/abc123/xyz":     {},
		"group/def456/abc/123": {},
		"user/jkl012/sss":      {},
	}

	invalidInput := make(map[string]SecretAlias, maxAliasCount+1)
	for i := 0; i <= maxAliasCount; i++ {
		invalidInput[fmt.Sprintf("agent/servant%d/test", i)] = SecretAlias{}
	}

	emptyInput := make(map[string]SecretAlias)

	if !IsValidAliases(validInput) {
		t.Errorf("Expected valid input, but got invalid")
	}

	if IsValidAliases(invalidInput) {
		t.Errorf("Expected invalid input, but got valid")
	}

	if !IsValidAliases(emptyInput) {
		t.Errorf("Expected valid input for empty map, but got invalid")
	}
}
