package encrypt

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"

	"github.com/pixconf/pixconf/internal/xid"
)

var (
	ErrWriteHash     = errors.New("error write hash")
	ErrWrongSecretID = errors.New("wrong secret id")
	ErrWrongVersion  = errors.New("wrong version")
)

func GetValueKey(epochKey []byte, secretID string, version int64) ([]byte, error) {
	hash := sha256.New()

	if !xid.IsValidSecretID(secretID) {
		return nil, ErrWrongSecretID
	}

	if len(epochKey) != KeySize {
		return nil, ErrKeySize
	}

	if version <= 0 {
		return nil, ErrWrongVersion
	}

	if _, err := hash.Write(epochKey); err != nil {
		return nil, ErrWriteHash
	}

	if _, err := hash.Write([]byte(secretID)); err != nil {
		return nil, ErrWriteHash
	}

	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, version)

	if _, err := hash.Write(buf[:n]); err != nil {
		return nil, ErrWriteHash
	}

	return hash.Sum(nil), nil
}
