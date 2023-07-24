package grpctools

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

func LoadTLSCredentials(certPath, keyPath string) (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
