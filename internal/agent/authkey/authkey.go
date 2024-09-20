package authkey

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"os"
)

type AuthKey struct {
	priv ed25519.PrivateKey
	pub  ed25519.PublicKey
}

func New(path string) (*AuthKey, error) {
	key := &AuthKey{}

	if len(path) > 1 {
		if stat, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
			if stat.IsDir() {
				return nil, errors.New("path is a directory")
			}

			persist, err := LoadFromDisk(path)
			if err != nil {
				return nil, err
			}

			if err := key.LoadKeys(persist.PrivateKey, persist.PublicKey); err != nil {
				return nil, err
			}
		} else {
			if err := key.generateKey(); err != nil {
				return nil, err
			}

			persist := &Persist{
				PrivateKey: key.priv,
				PublicKey:  key.pub,
			}

			if err := persist.SaveToDisk(path); err != nil {
				return nil, err
			}
		}

	} else {
		// generate ephemeral key
		if err := key.generateKey(); err != nil {
			return nil, err
		}
	}

	return key, nil
}

func (a *AuthKey) generateKey() error {
	var err error
	a.pub, a.priv, err = ed25519.GenerateKey(rand.Reader)

	return err
}

func (a *AuthKey) LoadKeys(private, public []byte) error {
	if len(private) != ed25519.PrivateKeySize {
		return errors.New("invalid private key size")
	}

	if len(public) != ed25519.PublicKeySize {
		return errors.New("invalid public key size")
	}

	a.priv = ed25519.PrivateKey(private)
	a.pub = a.priv.Public().(ed25519.PublicKey)

	return nil
}

func (a *AuthKey) Sign(data []byte) []byte {
	return ed25519.Sign(a.priv, data)
}
