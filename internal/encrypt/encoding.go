package encrypt

import "encoding/base64"

func EncodeToString(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func DecodeFromString(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
