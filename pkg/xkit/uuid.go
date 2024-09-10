package xkit

import "github.com/google/uuid"

// validate uuid or generate a new one
func GetUUID(input string) string {
	if uid, err := uuid.Parse(input); err == nil {
		return uid.String()
	}

	return uuid.New().String()
}
