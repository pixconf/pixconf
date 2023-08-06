package autocert

import (
	"encoding/pem"
)

func marshalCertificate(derBytes []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
}
