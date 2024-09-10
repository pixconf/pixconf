package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MQTT MQTTConfig `json:"mqtt"`
}

type MQTTConfig struct {
	Endpoints []string `json:"endpoints"`
}

func New() (*Config, error) {
	var conf Config

	if err := envconfig.Process("pixconf", &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
