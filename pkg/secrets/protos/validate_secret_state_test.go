package protos

import "testing"

func TestIsValidSecretState(t *testing.T) {
	validStates := []string{"normal", "hidden", "deleted"}
	invalidStates := []string{"invalid", "state", "names"}

	for _, state := range validStates {
		if !IsValidSecretState(state) {
			t.Errorf("Expected valid state for %s, but got invalid", state)
		}
	}

	for _, state := range invalidStates {
		if IsValidSecretState(state) {
			t.Errorf("Expected invalid state for %s, but got valid", state)
		}
	}
}

func TestSecretState_String(t *testing.T) {
	testCases := []struct {
		state    SecretState
		expected string
	}{
		{SecretStateNormal, "normal"},
		{SecretStateHidden, "hidden"},
		{SecretStateDeleted, "deleted"},
	}

	for _, testCase := range testCases {
		output := testCase.state.String()
		if output != testCase.expected {
			t.Errorf("For state %s, expected %s, but got %s", testCase.state, testCase.expected, output)
		}
	}
}
