package encrypt

import (
	"bytes"
	"encoding/base64"
	"testing"

	"golang.org/x/crypto/chacha20poly1305"
)

func TestChachaPoly_New(t *testing.T) {
	testCases := []struct {
		keySize       int
		expectedError bool
	}{
		{chacha20poly1305.KeySize, false},    // Valid key size
		{chacha20poly1305.KeySize - 1, true}, // Invalid key size
		{chacha20poly1305.KeySize + 1, true}, // Invalid key size
	}

	for _, tc := range testCases {
		key := make([]byte, tc.keySize)
		_, err := NewChachaPoly(key)

		if tc.expectedError && err == nil {
			t.Errorf("Expected error for key size %d, but got no error", tc.keySize)
		}
		if !tc.expectedError && err != nil {
			t.Errorf("Expected no error for key size %d, but got error: %s", tc.keySize, err)
		}
	}
}

func TestChachaPoly(t *testing.T) {
	encrtyptKey := "tmrSWtevJQ7nRZSLlMTNKrjpU10U9XX+McGRPK7hsHg="
	encrtyptData := []byte("the test message")

	encrtyptKeyBytes, err := base64.StdEncoding.DecodeString(encrtyptKey)
	if err != nil {
		t.Error(err)
	}

	enc, err := NewChachaPoly(encrtyptKeyBytes)
	if err != nil {
		t.Error(err)
	}

	if enc.GetEncryptType() != TypeChachaPoly {
		t.Error("wrong encrypt type")
	}

	testEncrypter(t, enc, encrtyptData)
	testEncrypterInvalidRand(t, enc, encrtyptData)

	encrypedData, err := enc.Encrypt(encrtyptData)
	if err != nil {
		t.Error(err)
	}

	dec, err := NewChachaPoly(encrtyptKeyBytes)
	if err != nil {
		t.Error(err)
	}

	decrypedData, err := dec.Decrypt(encrypedData)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(encrtyptData, decrypedData) {
		t.Error("wrong to decrypt")
	}
}
