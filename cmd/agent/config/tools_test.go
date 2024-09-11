package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnvOrDefault(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		def      string
		expected string
	}{
		{
			name:     "EnvVarSet",
			key:      "TEST_ENV_VAR",
			value:    "value",
			def:      "default",
			expected: "value",
		},
		{
			name:     "EnvVarNotSet",
			key:      "TEST_ENV_VAR",
			value:    "",
			def:      "default",
			expected: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != "" {
				os.Setenv(tt.key, tt.value)
			} else {
				os.Unsetenv(tt.key)
			}

			result := getEnvOrDefault(tt.key, tt.def)
			require.Equal(t, tt.expected, result)

			// Clean up
			os.Unsetenv(tt.key)
		})
	}
}

func TestGetEnvOrDefaultFunc(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		def      func() string
		expected string
	}{
		{
			name:     "EnvVarSet",
			key:      "TEST_ENV_VAR_FUNC",
			value:    "value",
			def:      func() string { return "default" },
			expected: "value",
		},
		{
			name:     "EnvVarNotSet",
			key:      "TEST_ENV_VAR_FUNC",
			value:    "",
			def:      func() string { return "default" },
			expected: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != "" {
				os.Setenv(tt.key, tt.value)
			} else {
				os.Unsetenv(tt.key)
			}

			result := getEnvOrDefaultFunc(tt.key, tt.def)
			require.Equal(t, tt.expected, result)

			// Clean up
			os.Unsetenv(tt.key)
		})
	}
}
