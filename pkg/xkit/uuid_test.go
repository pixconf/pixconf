package xkit

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUUID(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		overwrite bool
	}{
		{
			name:      "valid uuid",
			input:     "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			overwrite: false,
		},
		{
			name:      "invalid uuid",
			input:     "invalid",
			overwrite: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uid := GetUUID(tt.input)

			if tt.overwrite {
				_, err := uuid.Parse(uid)
				assert.Nil(t, err)

			} else {
				assert.Equal(t, tt.input, uid)
			}
		})
	}
}

func TestGetUUIDBytes(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		overwrite bool
	}{
		{
			name:      "valid uuid",
			input:     "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			overwrite: false,
		},
		{
			name:      "invalid uuid",
			input:     "invalid",
			overwrite: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uidBytes := GetUUIDBytes(tt.input)

			if tt.overwrite {
				_, err := uuid.FromBytes(uidBytes)
				assert.Nil(t, err)
			} else {
				expectedUUID, err := uuid.Parse(tt.input)
				assert.Nil(t, err)
				assert.Equal(t, expectedUUID[:], uidBytes)
			}
		})
	}
}
func TestLoadUUIDFromBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "valid uuid bytes",
			input:    []byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
			expected: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		},
		{
			name:     "invalid uuid bytes",
			input:    []byte{0x00, 0x01, 0x02},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LoadUUIDFromBytes(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
