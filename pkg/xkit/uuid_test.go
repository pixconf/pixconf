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
