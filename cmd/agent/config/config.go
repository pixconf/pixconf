package config

import (
	"errors"
	"fmt"
	"os"
)

func Load(configPath string) (*Config, error) {
	config := newConfig()

	fileInfo, err := os.Stat(configPath)
	if err == nil {
		// TODO: merge files in case of multiple config files
		if fileInfo.IsDir() {
			return nil, errors.New("config file is a directory")
		}

		filePayload, err := readFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}

		payload, err := parseConfigFile(configPath, filePayload)
		if err != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", err)
		}

		config.merge(payload)
	}

	return config, nil
}
