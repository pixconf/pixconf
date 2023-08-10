package encrypt

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/pixconf/pixconf/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestAesGCM_New(t *testing.T) {
	testCases := []struct {
		keySize       int
		expectedError bool
	}{
		{16, false}, // AES-128
		{24, false}, // AES-192
		{32, false}, // AES-256
		{15, true},  // Invalid key size
		{33, true},  // Invalid key size
	}

	for _, tc := range testCases {
		_, err := NewAesGCM(make([]byte, tc.keySize))

		if tc.expectedError && err == nil {
			t.Errorf("Expected error for key size %d, but got no error", tc.keySize)
		}
		if !tc.expectedError && err != nil {
			t.Errorf("Expected no error for key size %d, but got error: %s", tc.keySize, err)
		}
	}
}

func TestAesGCM(t *testing.T) {
	encrtyptKey := "tmrSWtevJQ7nRZSLlMTNKrjpU10U9XX+McGRPK7hsHg="
	encrtyptData := []byte("the test message")

	encrtyptKeyBytes, err := base64.StdEncoding.DecodeString(encrtyptKey)
	if err != nil {
		t.Error(err)
	}

	enc, err := NewAesGCM(encrtyptKeyBytes)
	if err != nil {
		t.Error(err)
	}

	if enc.GetEncryptType() != TypeAesGCM {
		t.Error("wrong encrypt type")
	}

	testEncrypter(t, enc, encrtyptData)
	testEncrypterInvalidRand(t, enc, encrtyptData)

	encrypedData, err := enc.Encrypt(encrtyptData)
	if err != nil {
		t.Error(err)
	}

	dec, err := NewAesGCM(encrtyptKeyBytes)
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

func testEncrypter(t *testing.T, enc Encrypter, encrtyptData []byte) {
	encrypedData, err := enc.Encrypt(encrtyptData)
	if err != nil {
		t.Error(err)
	}

	if encrypedData == nil {
		t.Error("encrypted data is nil")
	}

	encrypedDataSecond, err := enc.Encrypt(encrtyptData)
	if err != nil {
		t.Error(err)
	}

	if bytes.Equal(encrypedData, encrypedDataSecond) {
		t.Error("wrong generate nonce key")
	}

	if _, err := enc.Decrypt(encrypedData[:8]); err != ErrMalformed {
		t.Error(err)
	}

	decryptedData, err := enc.Decrypt(encrypedData)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(decryptedData, encrtyptData) {
		t.Error("wrong decrypd data")
	}
}

func testEncrypterInvalidRand(t *testing.T, enc Encrypter, encrtyptData []byte) {
	origRandReader := rand.Reader
	rand.Reader = mock.ErrorReader{}

	defer func() { rand.Reader = origRandReader }()

	encrypedData, err := enc.Encrypt(encrtyptData)

	if err == nil {
		t.Errorf("Encrypt should have returned an error")
	}

	assert.Empty(t, encrypedData, "Generated SecretID should be empty on error")
}
