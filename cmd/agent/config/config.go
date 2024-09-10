package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AgentAPISocket string `json:"agent_api_socket" yaml:"agent_api_socket"`
	AgentID        string `json:"agent_id" yaml:"agent_id"`
	Server         string `json:"server" yaml:"server"`
}

func Load(configPath string) (*Config, error) {
	var config *Config

	fileInfo, err := os.Stat(configPath)
	if err == nil {
		// TODO: merge files in case of multiple config files
		if fileInfo.IsDir() {
			return nil, errors.New("config file is a directory")
		}

		filePayload, err := os.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}

		switch filepath.Ext(configPath) {
		case ".json":
			if err := json.Unmarshal(filePayload, &config); err != nil {
				return nil, fmt.Errorf("failed to unmarshal json config: %w", err)
			}

		case ".yaml", ".yml":
			if err := yaml.Unmarshal(filePayload, &config); err != nil {
				return nil, fmt.Errorf("failed to unmarshal yaml config: %w", err)
			}

		default:
			return nil, errors.New("unsupported config file format")
		}
	}

	if err := defaults(config); err != nil {
		return nil, fmt.Errorf("failed to fill config defaults: %w", err)
	}

	return config, nil
}
