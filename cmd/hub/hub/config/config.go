package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
}

func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("hub", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
