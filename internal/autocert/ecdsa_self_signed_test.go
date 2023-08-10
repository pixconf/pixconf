package autocert

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestGenerateSelfSignedECDSACert(t *testing.T) {
	certPEM, keyPEM, err := GenerateSelfSignedECDSACert("test_instance")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if certPEM == nil {
		t.Error("Certificate can not be empty")
	}

	if keyPEM == nil {
		t.Error("Certificate key can not be empty")
	}

	certBlock, _ := pem.Decode(certPEM)
	if certBlock == nil {
		t.Fatal("Failed to decode certificate PEM")
	}

	certificate, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		t.Errorf("Failed to parse certificate pair: %s", err)
	}

	if len(certificate.DNSNames) != 3 {
		t.Errorf("Expected 3 DNSNames in certificate, but got %d", len(certificate.DNSNames))
	}

	if len(certificate.IPAddresses) != 2 {
		t.Errorf("Expected 2 IPAddresses in certificate, but got %d", len(certificate.IPAddresses))
	}

	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil {
		t.Fatal("Failed to decode private key PEM")
	}

	privateKey, err := x509.ParseECPrivateKey(keyBlock.Bytes)
	if err != nil {
		t.Errorf("Failed to parse private key: %s", err)
	}

	if privateKey == nil {
		t.Errorf("Private key can not be empty")
	}
}
