package protos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecretCreateRequest_Validate(t *testing.T) {
	tests := []struct {
		name          string
		request       SecretCreateRequest
		expectedError bool
	}{
		{
			name: "ValidRequest",
			request: SecretCreateRequest{
				Description: "Test secret",
				State:       SecretStateNormal.String(),
				Tags:        []string{"tag1", "tag2"},
				Alias: map[string]SecretAlias{
					"agent/thename/thekey": {},
				},
			},
			expectedError: false,
		},
		{
			name: "InvalidTags",
			request: SecretCreateRequest{
				Description: "Invalid tags",
				State:       SecretStateNormal.String(),
				Tags:        []string{"tag1", "@tag2"}, // Invalid tag
			},
			expectedError: true,
		},
		{
			name: "InvalidAlias",
			request: SecretCreateRequest{
				Description: "Invalid alias",
				State:       SecretStateNormal.String(),
				Tags:        []string{"tag1", "tag2"},
				Alias: map[string]SecretAlias{
					"alias1": {},
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errors := test.request.Validate()

			if test.expectedError {
				assert.NotEmpty(t, errors, "Expected validation errors, but got none")
			} else {
				assert.Empty(t, errors, "Expected no validation errors, but got some")
			}
		})
	}
}
