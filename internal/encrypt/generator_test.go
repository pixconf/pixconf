package encrypt

import "testing"

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Error(err)
	}

	if len(key) != KeySize {
		t.Error("error key size")
	}
}
