package autocert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"
)

const guestDomain = "guest.pixconf.vitalvas.dev"

func GenerateSelfSignedECDSACert(instance string) ([]byte, []byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	notBefore := time.Now().Add(-time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	domain := fmt.Sprintf("pcss-%s.%s", strconv.FormatInt(notBefore.UnixMicro(), 16), guestDomain)

	template := x509.Certificate{
		SerialNumber: serialNumber,

		Subject: pkix.Name{
			CommonName:         domain,
			Organization:       []string{"PixConf Automation"},
			OrganizationalUnit: []string{"PixConf Self Signed"},
		},

		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(365 * 24 * time.Hour), // Valid for 1 year
		IsCA:                  true,
		BasicConstraintsValid: true,

		DNSNames: []string{
			"localhost",
			domain,
			fmt.Sprintf("%s.%s", instance, domain),
		},
		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
			net.ParseIP("::1"),
		},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	certPEM := marshalCertificate(derBytes)
	keyPEM, err := marshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}

	return certPEM, keyPEM, nil
}
