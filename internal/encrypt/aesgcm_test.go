package encrypt

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestAesGCM(t *testing.T) {
	encrtyptKey := "tmrSWtevJQ7nRZSLlMTNKrjpU10U9XX+McGRPK7hsHg="

	encrtyptKeyBytes, err := base64.StdEncoding.DecodeString(encrtyptKey)
	if err != nil {
		t.Error(err)
	}

	enc, err := NewAesGCM(encrtyptKeyBytes)
	if err != nil {
		t.Error(err)
	}

	testEncrypter(t, enc)
}

func testEncrypter(t *testing.T, enc Encrypter) {
	encrtyptData := []byte("the test message")

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
