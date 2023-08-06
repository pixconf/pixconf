package autocert

import "crypto/tls"

func GetTLSConfig(cert, privateKey []byte) (*tls.Config, error) {
	certificate, err := tls.X509KeyPair(cert, privateKey)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}

	return config, nil
}
