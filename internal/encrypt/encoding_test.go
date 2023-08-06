package encrypt

import (
	"bytes"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	data := []byte{0xa, 0xb, 0xc, 0xd, 0xe, 0xf}

	encoded := EncodeToString(data)

	decoded, err := DecodeFromString(encoded)

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, decoded) {
		t.Error("error encode-decode")
	}
}
