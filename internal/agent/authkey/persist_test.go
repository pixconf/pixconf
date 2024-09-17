package authkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersistMarshal(t *testing.T) {
	tests := []struct {
		name    string
		persist Persist
		want    string
		wantErr bool
	}{
		{
			name: "valid keys",
			persist: Persist{
				PrivateKey: []byte("private_key"),
				PublicKey:  []byte("public_key"),
			},
			want:    "eyJwcml2YXRlIjoiY0hKcGRtRjBaVjlyWlhrPSIsInB1YmxpYyI6ImNIVmliR2xqWDJ0bGVRPT0ifQ==",
			wantErr: false,
		},
		{
			name: "empty keys",
			persist: Persist{
				PrivateKey: nil,
				PublicKey:  nil,
			},
			want:    "e30=",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.persist.Marshal()
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
