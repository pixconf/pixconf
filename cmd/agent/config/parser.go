package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func readFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func parseConfigFile(filePath string, payload []byte, config *Config) error {
	switch filepath.Ext(filePath) {
	case ".json":
		decoder := json.NewDecoder(bytes.NewReader(payload))
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&config); err != nil {
			return fmt.Errorf("failed to unmarshal json config: %w", err)
		}

		return nil

	case ".yaml", ".yml":
		decoder := yaml.NewDecoder(bytes.NewReader(payload))
		decoder.KnownFields(true)

		if err := decoder.Decode(&config); err != nil {
			return fmt.Errorf("failed to unmarshal yaml config: %w", err)
		}

		return nil
	}

	return fmt.Errorf("unsupported config file format: %s", filepath.Ext(filePath))
}