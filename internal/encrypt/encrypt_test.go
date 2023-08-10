package encrypt

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypter(t *testing.T) {
	encrtyptKey := "tmrSWtevJQ7nRZSLlMTNKrjpU10U9XX+McGRPK7hsHg="

	enc, err := NewEncoded(encrtyptKey, TypeAesGCM)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(enc) != reflect.TypeOf(&AesGCM{}) {
		t.Error("error open AES-GCM")
	}

	enc, err = NewEncoded(encrtyptKey, TypeChachaPoly)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(enc) != reflect.TypeOf(&ChachaPoly{}) {
		t.Error("error open Chacha-Poly")
	}

	if _, err = NewEncoded(encrtyptKey, 99); err != ErrUnknownEncryptType {
		t.Error(err)
	}
}

func TestNew(t *testing.T) {
	_, err := New([]byte{0x0, 0x1, 0x2}, TypeAesGCM)
	if err != ErrKeySize {
		t.Error("error check key size")
	}
}

func TestNewEncoded_Valid(t *testing.T) {
	enc, err := NewEncoded("HXI6QViGe0itwp7tq4F+0awGOIENNqoj1vMoUgBCQfg=", TypeAesGCM)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(enc) != reflect.TypeOf(&AesGCM{}) {
		t.Error("error open AES-GCM")
	}
}

func TestNewEncoded_Invalid(t *testing.T) {
	enc, err := NewEncoded("asd", TypeAesGCM)

	if err == nil {
		t.Errorf("NewEncoded should have returned an error")
	}

	assert.Empty(t, enc, "NewEncoded Key should be empty on error")
}
