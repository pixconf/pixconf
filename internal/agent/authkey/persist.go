package authkey

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	Base64PersistEncoding = base64.StdEncoding
)

type Persist struct {
	PrivateKey []byte `json:"private,omitempty"`
	PublicKey  []byte `json:"public,omitempty"`
}

func (p *Persist) Marshal() (string, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return Base64PersistEncoding.EncodeToString(bytes), nil
}

func LoadFromDisk(path string) (*Persist, error) {
	filePayload, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	payloadDecoded, err := Base64PersistEncoding.DecodeString(string(filePayload))
	if err != nil {
		return nil, err
	}

	persist := &Persist{}
	if err := json.Unmarshal(payloadDecoded, &persist); err != nil {
		return nil, err
	}

	return persist, nil
}

func (p *Persist) SaveToDisk(path string) error {
	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	encoded := Base64PersistEncoding.EncodeToString(bytes)

	keyDir := filepath.Dir(path)
	if _, err := os.Stat(keyDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(keyDir, 0750); err != nil {
			return err
		}
	}

	return os.WriteFile(path, []byte(encoded), 0600)
}
