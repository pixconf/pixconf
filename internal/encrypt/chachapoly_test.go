package encrypt

import (
	"encoding/base64"
	"testing"
)

func TestChachaPoly(t *testing.T) {
	encrtyptKey := "tmrSWtevJQ7nRZSLlMTNKrjpU10U9XX+McGRPK7hsHg="

	encrtyptKeyBytes, err := base64.StdEncoding.DecodeString(encrtyptKey)
	if err != nil {
		t.Error(err)
	}

	enc, err := NewChachaPoly(encrtyptKeyBytes)
	if err != nil {
		t.Error(err)
	}

	testEncrypter(t, enc)
}
