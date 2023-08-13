package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	APIAddress  string `split_words:"true" default:"[::]:8140"`
	TLSCertPath string `split_words:"true"`
	TLSKeyPath  string `split_words:"true"`
}

func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("hub", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
