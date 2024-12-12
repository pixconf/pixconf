package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MQTT MQTTConfig `split_words:"true"`
}

type MQTTConfig struct {
	Listen    []string `split_words:"true"`
	Endpoints []string `split_words:"true"`
}

func New() (*Config, error) {
	var conf Config

	if err := envconfig.Process("pixconf", &conf); err != nil {
		return nil, err
	}

	if conf.MQTT.Listen == nil {
		conf.MQTT.Listen = []string{"mqtt://:1883"}
	}

	if err := validate(conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
