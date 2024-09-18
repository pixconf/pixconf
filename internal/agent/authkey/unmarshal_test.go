package authkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalHeader(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *AuthHeader
		expectedErr bool
	}{
		{
			name:        "ValidHeader",
			input:       `eyJwayI6ICJhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5eiJ9`, // base64 for {"pk": "abcdefghijklmnopqrstuvwxyz"}
			expected:    &AuthHeader{PublicKey: "abcdefghijklmnopqrstuvwxyz"},
			expectedErr: false,
		},
		{
			name:        "InvalidBase64",
			input:       `!!!invalidbase64!!!`,
			expected:    nil,
			expectedErr: true,
		},
		{
			name:        "InvalidJSON",
			input:       `eyJwayI6ICJhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5eiI`, // base64 for {"pk": "abcdefghijklmnopqrstuvwxyz" (missing closing brace)
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, err := unmarshalHeader(tt.input)
			if tt.expectedErr && err == nil {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.expected, header)
		})
	}
}

func TestUnmarshalPayload(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *AuthPayload
		expectedErr bool
	}{
		{
			name:        "ValidPayload",
			input:       `eyJpc3MiOiAiZXhhbXBsZS01Njg1YmRiODU5LXhtbWdkIn0`, // base64 for {"iss": "example-5685bdb859-xmmgd"}
			expected:    &AuthPayload{Issuer: "example-5685bdb859-xmmgd"},
			expectedErr: false,
		},
		{
			name:        "InvalidBase64",
			input:       `!!!invalidbase64!!!`,
			expected:    nil,
			expectedErr: true,
		},
		{
			name:        "InvalidJSON",
			input:       `eyJpc3MiOiAiZXhhbXBsZS01Njg1YmRiODU5LXhtbWdkIg`, // base64 for {"iss": "example-5685bdb859-xmmgd" (missing closing brace)
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := unmarshalPayload(tt.input)
			if tt.expectedErr && err == nil {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.expected, payload)
		})
	}
}
