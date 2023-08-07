package encrypt

import (
	"bytes"
	"encoding/base64"
	"testing"
)

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
