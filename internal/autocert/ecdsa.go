package autocert

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

func marshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: der}), nil
}
