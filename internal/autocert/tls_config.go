package autocert

import "crypto/tls"

func GetTLSConfig(cert, privateKey []byte) (*tls.Config, error) {
	certificate, err := tls.X509KeyPair(cert, privateKey)
	if err != nil {
		return nil, err
	}

	// https://ssl-config.mozilla.org/
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},
		Certificates: []tls.Certificate{
			certificate,
		},
	}

	return config, nil
}
