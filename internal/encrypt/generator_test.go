package encrypt

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pixconf/pixconf/internal/mock"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Error(err)
	}

	if len(key) != KeySize {
		t.Error("error key size")
	}
}

func TestGenerateKey_Error(t *testing.T) {
	rand.Reader = mock.ErrorReader{}

	secretID, err := GenerateKey()
	if err == nil {
		t.Errorf("GenerateKey should have returned an error")
	}

	assert.Empty(t, secretID, "Generated Key should be empty on error")
}
