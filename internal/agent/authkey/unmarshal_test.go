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

func TestUnmarshalCheckSign(t *testing.T) {
	const (
		privateKeyBase64 = "dBYe_mshXFvYwYpwfC41OeHUncz0N0y7HcFymJhA0hae_19QeQqnFQEgpu2WmPgwph4sBaehV_zlqnwYcnxwyw"
		publicKeyBase64  = "nv9fUHkKpxUBIKbtlpj4MKYeLAWnoVf85ap8GHJ8cMs"
	)

	tests := []struct {
		name        string
		pubKey      string
		signedData  []byte
		signRaw     []byte
		expectedErr bool
	}{
		{
			name:        "ValidSignature",
			pubKey:      publicKeyBase64,
			signedData:  []byte("test data"),
			signRaw:     []byte{0xa6, 0x7a, 0x97, 0x6f, 0xca, 0x9a, 0xe3, 0x60, 0x43, 0x30, 0x25, 0xf6, 0x0, 0x7d, 0xcd, 0xdc, 0x91, 0x40, 0x74, 0x27, 0x44, 0xcd, 0xb4, 0x2, 0xec, 0x9e, 0xfc, 0xf1, 0x41, 0xfc, 0x85, 0xb3, 0x2f, 0xf7, 0x5f, 0x50, 0xa4, 0x51, 0xf5, 0xf1, 0x8d, 0xd4, 0xaf, 0x56, 0x3e, 0x60, 0xd0, 0xe6, 0x8f, 0x3d, 0x52, 0x36, 0xf0, 0x76, 0x91, 0x1f, 0x8a, 0x84, 0x36, 0x43, 0x0, 0x36, 0xed, 0xb},
			expectedErr: false,
		},
		{
			name:        "InvalidPublicKey",
			pubKey:      "invalidbase64",
			signedData:  nil,
			signRaw:     nil,
			expectedErr: true,
		},
		{
			name:        "InvalidPublicKeySize",
			pubKey:      "nv9fUHkKpxUBIKbt",
			signedData:  nil,
			signRaw:     nil,
			expectedErr: true,
		},
		{
			name:        "InvalidSignature",
			pubKey:      publicKeyBase64,
			signedData:  []byte("test data"),
			signRaw:     []byte{0xFF, 0xFF, 0xFF, 0xFF},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := unmarshalCheckSign(tt.pubKey, tt.signedData, tt.signRaw)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
