package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	defer os.Remove(tmpFile.Name())

	content := []byte("test content")
	if _, err := tmpFile.Write(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	tests := []struct {
		name     string
		filePath string
		want     []byte
		wantErr  bool
	}{
		{"ValidFile", tmpFile.Name(), content, false},
		{"NonExistentFile", "nonexistentfile.txt", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.filePath)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseConfigFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		payload  []byte
		want     Config
		wantErr  bool
	}{
		{"ValidJSON", "config.json", []byte(`{"agent_id": "value1"}`), Config{AgentID: "value1"}, false},
		{"ValidYAML", "config.yaml", []byte("agent_id: value1\n"), Config{AgentID: "value1"}, false},
		{"ValidYAMLShort", "config.yml", []byte("agent_id: value1"), Config{AgentID: "value1"}, false},
		{"InvalidContent", "config.json", []byte("invalid content"), Config{}, true},
		{"EmptyFileJSON", "config.json", []byte(""), Config{}, true},
		{"EmptyFileYAML", "config.yaml", []byte(""), Config{}, true},
		{"UnsupportedFormat", "config.txt", []byte("agent_id: value1\n"), Config{}, true},
		{"NoFileExtension", "config", []byte("agent_id: value1\n"), Config{}, true},
		{"InvalidJSON", "config.json", []byte(`{"agent_id": 123}`), Config{}, true},
		{"InvalidYAML", "config.yaml", []byte("agent_id: ***1\n"), Config{}, true},
		{"UnknownFieldJSON", "config.json", []byte(`{"unknown_fileld_test": "value1"}`), Config{}, true},
		{"UnknownFieldYAML", "config.yaml", []byte("unknown_fileld_test: value1\n"), Config{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfigFile(tt.filePath, tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && *got != tt.want {
				t.Errorf("parseConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
