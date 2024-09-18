package authkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	privateKeyBase64 = "3niPGZ21km9xR665ML9cYJdbLfTqaG6hfxiYhatmUx/q0LbOvyaG0HehbSoJR1DWjVVMfrw5cVRJO3nJDWGWnw=="
	publicKeyBase64  = "6tC2zr8mhtB3oW0qCUdQ1o1VTH68OXFUSTt5yQ1hlp8="
)

func TestLoadKeys(t *testing.T) {
	validPubKey, err := Base64PersistEncoding.DecodeString(publicKeyBase64)
	assert.Nil(t, err)

	validPrivKey, err := Base64PersistEncoding.DecodeString(privateKeyBase64)
	assert.Nil(t, err)

	tests := []struct {
		name    string
		pubKey  []byte
		privKey []byte
		wantErr bool
	}{
		{
			name:    "Valid public and private key",
			pubKey:  validPubKey,
			privKey: validPrivKey,
			wantErr: false,
		},
		{
			name:    "Invalid private key",
			pubKey:  validPubKey,
			privKey: []byte(""),
			wantErr: true,
		},
		{
			name:    "Invalid public key",
			pubKey:  []byte(""),
			privKey: validPrivKey,
			wantErr: true,
		},
		{
			name:    "Invalid public and private key",
			pubKey:  []byte("=="),
			privKey: []byte("=="),
			wantErr: true,
		},
		{
			name:    "Empty public and private key",
			pubKey:  nil,
			privKey: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &AuthKey{}

			err := key.LoadKeys(tt.privKey, tt.pubKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
