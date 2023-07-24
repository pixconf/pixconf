package encrypt

import (
	"crypto/sha256"
	"encoding/binary"
)

func GetValueKey(epochKey []byte, secretID string, version int64) ([]byte, error) {
	hash := sha256.New()

	if _, err := hash.Write(epochKey); err != nil {
		return nil, err
	}

	if _, err := hash.Write([]byte(secretID)); err != nil {
		return nil, err
	}

	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, version)

	if _, err := hash.Write(buf[:n]); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
