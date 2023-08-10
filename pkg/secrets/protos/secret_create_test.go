package protos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecretCreateRequest_Validate(t *testing.T) {
	invalidCountTags := make([]string, 300)
	invalidCountAlias := make(map[string]SecretAlias, 300)

	for x := 0; x < 300; x++ {
		invalidCountTags[x] = fmt.Sprintf("tag%d", x)
		invalidCountAlias[fmt.Sprintf("alias/%d", x)] = SecretAlias{}
	}

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
			name: "InvalidState",
			request: SecretCreateRequest{
				State: "test",
			},
			expectedError: true,
		},
		{
			name: "InvalidTag",
			request: SecretCreateRequest{
				Description: "Invalid tags",
				State:       SecretStateNormal.String(),
				Tags:        []string{"tag1", "@tag2"}, // Invalid tag
			},
			expectedError: true,
		},
		{
			name: "InvalidTags",
			request: SecretCreateRequest{
				State: SecretStateNormal.String(),
				Tags:  invalidCountTags,
			},
			expectedError: true,
		},
		{
			name: "InvalidAliases",
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
		{
			name: "InvalidAlias",
			request: SecretCreateRequest{
				State: SecretStateNormal.String(),
				Alias: invalidCountAlias,
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
