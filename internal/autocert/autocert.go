package autocert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

const guestDomain = "guest.pixconf.vitalvas.dev"

func MakeSelfSigned(apps []string, instance string) error {
	notBefore := time.Now()
	notAfter := notBefore.Add(13 * 30 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"PixConf Self Signed"},
		},

		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage: x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}

	hosts := []string{"localhost", "127.0.0.1", "::1"}
	for _, row := range apps {
		hosts = append(hosts, fmt.Sprintf("%s.pcss-%s.%s", row, instance, guestDomain))
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	priv, err := makePrivateKey(instance)
	if err != nil {
		return err
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	// write cert
	certOut, err := os.Create(fmt.Sprintf("tmp/%s-cert.pem", instance))
	if err != nil {
		return err
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return err
	}

	return certOut.Close()
}

func makePrivateKey(instance string) (*ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	keyOut, err := os.OpenFile(fmt.Sprintf("tmp/%s-key.pem", instance), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}

	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return nil, err
	}

	return priv, keyOut.Close()
}
