package authkey

import (
	"crypto/ed25519"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	privateKeyBase64 = "3niPGZ21km9xR665ML9cYJdbLfTqaG6hfxiYhatmUx/q0LbOvyaG0HehbSoJR1DWjVVMfrw5cVRJO3nJDWGWnw=="
	publicKeyBase64  = "6tC2zr8mhtB3oW0qCUdQ1o1VTH68OXFUSTt5yQ1hlp8="
)

func TestLoadPrivateKey(t *testing.T) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	assert.Nil(t, err)

	assert.True(t, len(privateKeyBytes) == ed25519.PrivateKeySize)

	expectedPubKey, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	assert.Nil(t, err)

	assert.True(t, len(expectedPubKey) == ed25519.PublicKeySize)

	expectedPubKeyValid := ed25519.PublicKey(expectedPubKey)

	authKey := &AuthKey{}
	err = authKey.LoadPrivateKey(privateKeyBytes)
	assert.Nil(t, err)

	assert.NotNil(t, authKey.priv)
	assert.NotNil(t, authKey.priv)

	signed := authKey.Sign([]byte("test"))
	assert.NotNil(t, signed)

	assert.True(t, ed25519.Verify(expectedPubKeyValid, []byte("test"), signed))

	assert.Equal(t, expectedPubKeyValid, authKey.pub)
}

func TestLoadPrivateKeyInvalid(t *testing.T) {
	authKey := &AuthKey{}

	err := authKey.LoadPrivateKey([]byte("invalid"))
	assert.NotNil(t, err)
}
