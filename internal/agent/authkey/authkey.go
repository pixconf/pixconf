package authkey

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
)

type AuthKey struct {
	priv ed25519.PrivateKey
	pub  ed25519.PublicKey
}

func New() *AuthKey {
	key := &AuthKey{}

	if err := key.generateKey(); err != nil {
		return nil
	}

	return key
}

func (a *AuthKey) generateKey() error {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	a.pub = publicKey
	a.priv = privateKey

	return nil
}

func (a *AuthKey) LoadPrivateKey(data []byte) error {
	if len(data) != ed25519.PrivateKeySize {
		return errors.New("invalid private key size")
	}

	a.priv = ed25519.PrivateKey(data)
	a.pub = a.priv.Public().(ed25519.PublicKey)

	return nil
}

func (a *AuthKey) Sign(data []byte) []byte {
	return ed25519.Sign(a.priv, data)
}
