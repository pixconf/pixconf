package authkey

import (
	"encoding/json"
	"os"
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

func TestLoadFromDisk(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		want        *Persist
		wantErr     bool
	}{
		{
			name:        "valid file",
			fileContent: "eyJwcml2YXRlIjoiY0hKcGRtRjBaVjlyWlhrPSIsInB1YmxpYyI6ImNIVmliR2xqWDJ0bGVRPT0ifQ==",
			want: &Persist{
				PrivateKey: []byte("private_key"),
				PublicKey:  []byte("public_key"),
			},
			wantErr: false,
		},
		{
			name:        "invalid base64",
			fileContent: "invalid_base64",
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "invalid json",
			fileContent: "aW52YWxpZF9qc29u", // base64 for "invalid_json"
			want:        nil,
			wantErr:     true,
		},
		{
			name:    "non-existent file",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.name != "non-existent file" {
				tmpFile, err := os.CreateTemp("", "testfile")
				assert.NoError(t, err)
				defer os.Remove(tmpFile.Name())

				_, err = tmpFile.WriteString(tt.fileContent)
				assert.NoError(t, err)
				tmpFile.Close()

				filePath = tmpFile.Name()

			} else {
				filePath = "non_existent_file"
			}

			got, err := LoadFromDisk(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromDisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSaveToDisk(t *testing.T) {
	tests := []struct {
		name    string
		persist Persist
		wantErr bool
	}{
		{
			name: "valid persist data",
			persist: Persist{
				PrivateKey: []byte("private_key"),
				PublicKey:  []byte("public_key"),
			},
			wantErr: false,
		},
		{
			name: "invalid directory path",
			persist: Persist{
				PrivateKey: []byte("private_key"),
				PublicKey:  []byte("public_key"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.name == "invalid directory path" {
				filePath = "/invalid_dir/testfile"
			} else {
				tmpFile, err := os.CreateTemp("", "testfile")
				assert.NoError(t, err)
				defer os.Remove(tmpFile.Name())
				filePath = tmpFile.Name()
			}

			err := tt.persist.SaveToDisk(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveToDisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify file content
				fileContent, err := os.ReadFile(filePath)
				assert.NoError(t, err)

				decodedContent, err := Base64PersistEncoding.DecodeString(string(fileContent))
				assert.NoError(t, err)

				var got Persist
				err = json.Unmarshal(decodedContent, &got)
				assert.NoError(t, err)

				assert.Equal(t, tt.persist, got)
			}
		})
	}
}
