package autocert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestMarshalECPrivateKey(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Error generating ECDSA private key: %v", err)
	}

	keyPEM, err := marshalECPrivateKey(privateKey)
	if err != nil {
		t.Fatalf("Error marshaling ECDSA private key: %v", err)
	}

	block, _ := pem.Decode(keyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		t.Fatalf("Failed to decode PEM block")
	}

	parsedPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		t.Fatal(err)
	}

	if parsedPrivateKey.D.Cmp(privateKey.D) != 0 {
		t.Errorf("Parsed ECDSA private key does not match original")
	}
}
