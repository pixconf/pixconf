package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	AgentAPISocket string `json:"agent_api_socket" yaml:"agent_api_socket"`
	AgentID        string `json:"agent_id" yaml:"agent_id"`
	Server         string `json:"server" yaml:"server"`
}

func Load(configPath string) (*Config, error) {
	config := &Config{}

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

		if err := parseConfigFile(configPath, filePayload, config); err != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", err)
		}
	}

	if err := defaults(config); err != nil {
		return nil, fmt.Errorf("failed to fill config defaults: %w", err)
	}

	return config, nil
}
