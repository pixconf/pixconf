package authkey

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/pixconf/pixconf/internal/buildinfo"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAuthHeader(t *testing.T) {
	authKey, err := New("")
	assert.Nil(t, err)

	err = authKey.generateKey()
	assert.Nil(t, err)

	encodedHeader, err := authKey.generateAuthHeader()
	assert.Nil(t, err)

	decodedHeader, err := base64.RawURLEncoding.DecodeString(encodedHeader)
	assert.Nil(t, err)

	var header AuthHeader
	err = json.Unmarshal(decodedHeader, &header)
	assert.Nil(t, err)

	expectedAlgorithm := "ed25519"
	expectedPublicKey := base64.RawURLEncoding.EncodeToString(authKey.pub)

	assert.Equal(t, expectedAlgorithm, header.Algorithm)
	assert.Equal(t, expectedPublicKey, header.PublicKey)
}

func TestGenerateAuthPayload(t *testing.T) {
	mockAgentID := "mockAgentID"
	authKey, err := New("")
	assert.Nil(t, err)

	err = authKey.generateKey()
	assert.Nil(t, err)

	encodedPayload, err := authKey.generateAuthPayload(mockAgentID)
	assert.Nil(t, err)

	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	assert.Nil(t, err)

	var payload AuthPayload
	err = json.Unmarshal(decodedPayload, &payload)
	assert.Nil(t, err)

	assert.Equal(t, mockAgentID, payload.Issuer)
	assert.NotEmpty(t, payload.JwtID)
	assert.NotZero(t, payload.IssuedAT)
	assert.Equal(t, buildinfo.Version, payload.Version)

	_, err = uuid.Parse(payload.JwtID)
	assert.Nil(t, err)

	assert.True(t, payload.IssuedAT > 0)
}

func TestGenerateAuthKey(t *testing.T) {
	mockAgentID := "mockAgentID"

	authKey, err := New("")
	assert.Nil(t, err)

	err = authKey.generateKey()
	assert.Nil(t, err)

	encodedAuthKey, err := authKey.GenerateAuthKey(mockAgentID)
	assert.Nil(t, err)

	assert.True(t, len(encodedAuthKey) > 0)
	assert.True(t, len(mockAgentID)+len(encodedAuthKey) < 65000)

	parts := strings.Split(string(encodedAuthKey), ".")
	assert.Len(t, parts, 3)

	// Validate Header
	decodedHeader, err := base64.RawURLEncoding.DecodeString(parts[0])
	assert.Nil(t, err)

	var header AuthHeader
	err = json.Unmarshal(decodedHeader, &header)
	assert.Nil(t, err)

	expectedAlgorithm := "ed25519"
	expectedPublicKey := base64.RawURLEncoding.EncodeToString(authKey.pub)

	assert.Equal(t, expectedAlgorithm, header.Algorithm)
	assert.Equal(t, expectedPublicKey, header.PublicKey)

	// Validate Payload
	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	assert.Nil(t, err)

	var payload AuthPayload
	err = json.Unmarshal(decodedPayload, &payload)
	assert.Nil(t, err)

	assert.Equal(t, mockAgentID, payload.Issuer)
	assert.NotEmpty(t, payload.JwtID)
	assert.NotZero(t, payload.IssuedAT)

	_, err = uuid.Parse(payload.JwtID)
	assert.Nil(t, err)

	assert.True(t, payload.IssuedAT > 0)

	// Validate Signature
	dataToSign := fmt.Sprintf("%s.%s", parts[0], parts[1])
	expectedSignature := authKey.Sign([]byte(dataToSign))
	encodedExpectedSignature := base64.RawURLEncoding.EncodeToString(expectedSignature)

	assert.Equal(t, encodedExpectedSignature, parts[2])
}

func TestGenerateAuthKeyLen(t *testing.T) {
	tests := []struct {
		name        string
		agentID     string
		expectError bool
	}{
		{
			name:        "short agentID",
			agentID:     "a",
			expectError: false,
		},
		{
			name:        "long agentID",
			agentID:     strings.Repeat("a", 1024),
			expectError: false,
		},
		{
			name:        "long 4k agentID",
			agentID:     strings.Repeat("a", 4096),
			expectError: false,
		},
		{
			name:        "overlong agentID",
			agentID:     strings.Repeat("a", 65535),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authKey, err := New("")
			assert.Nil(t, err)

			err = authKey.generateKey()
			assert.Nil(t, err)

			encodedAuthKey, err := authKey.GenerateAuthKey(tt.agentID)
			assert.Nil(t, err)

			assert.Nil(t, err)
			assert.True(t, len(encodedAuthKey) > 0)
			if tt.expectError {
				assert.True(t, len(tt.agentID)+len(encodedAuthKey) >= 65535)
			} else {
				assert.True(t, len(tt.agentID)+len(encodedAuthKey) < 65535)
			}
		})
	}
}
