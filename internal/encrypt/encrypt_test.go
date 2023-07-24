package encrypt

import (
	"reflect"
	"testing"
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
