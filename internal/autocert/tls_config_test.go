package autocert

import (
	"errors"
	"testing"
)

func TestGetTLSConfig(t *testing.T) {
	validCert, validPrivateKey, err := GenerateSelfSignedECDSACert("test")
	if err != nil {
		t.Error(err)
	}

	invalidCert := []byte("invalid_cert")
	invalidPrivateKey := []byte("invalid_private_key")

	config, err := GetTLSConfig(validCert, validPrivateKey)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(config.Certificates) != 1 {
		t.Errorf("Expected 1 certificate, but got %d", len(config.Certificates))
	}

	if _, err = GetTLSConfig(invalidCert, validPrivateKey); err == nil {
		t.Error("Expected error due to invalid certificate, but got nil")
	}

	if _, err = GetTLSConfig(validCert, invalidPrivateKey); err == nil {
		t.Error("Expected error due to invalid private key, but got nil")
	}

	if _, err = GetTLSConfig(invalidCert, invalidPrivateKey); err == nil {
		t.Error("Expected error due to invalid certificate and private key, but got nil")
	}

	customErr := errors.New("custom error")
	_, err = GetTLSConfig(validCert, validPrivateKey)
	if err != nil && err != customErr {
		t.Errorf("Expected custom error, but got: %s", err)
	}
}
