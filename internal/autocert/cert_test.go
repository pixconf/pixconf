package autocert

import (
	"bytes"
	"encoding/pem"
	"testing"
)

func TestMarshalCertificate(t *testing.T) {
	derBytes := []byte{0x30, 0x82, 0x01, 0x23}

	expectedPemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}

	pemBytes := marshalCertificate(derBytes)
	decodedPemBlock, _ := pem.Decode(pemBytes)

	if decodedPemBlock == nil {
		t.Errorf("Failed to decode PEM block")
		return
	}

	if decodedPemBlock.Type != expectedPemBlock.Type {
		t.Errorf("Expected PEM block type %s, but got %s", expectedPemBlock.Type, decodedPemBlock.Type)
	}

	if !bytes.Equal(decodedPemBlock.Bytes, expectedPemBlock.Bytes) {
		t.Error("Decoded PEM block bytes do not match expected bytes")
	}
}
