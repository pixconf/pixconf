package xkit

import "github.com/google/uuid"

// validate uuid or generate a new one
func GetUUID(input string) string {
	if uid, err := uuid.Parse(input); err == nil {
		return uid.String()
	}

	return uuid.New().String()
}

// response 16 bytes of uuid
func GetUUIDBytes(input string) []byte {
	if uid, err := uuid.Parse(input); err == nil {
		return uid[:]
	}

	uid := uuid.New()
	return uid[:]
}

func LoadUUIDFromBytes(input []byte) string {
	if uid, err := uuid.ParseBytes(input); err == nil {
		return uid.String()
	}

	return ""
}
