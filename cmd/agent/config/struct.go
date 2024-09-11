package config

import "reflect"

type Config struct {
	AgentAPISocket string `json:"agent_api_socket" yaml:"agent_api_socket"`
	AgentID        string `json:"agent_id" yaml:"agent_id"`
	Server         string `json:"server" yaml:"server"`
}

func (c *Config) merge(customConfig *Config) {
	defaultValue := reflect.ValueOf(c).Elem()
	customValue := reflect.ValueOf(customConfig).Elem()

	for i := 0; i < customValue.NumField(); i++ {
		field := customValue.Field(i)
		if !isZeroValue(field) {
			defaultValue.Field(i).Set(field)
		}
	}
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func newConfig() *Config {
	return &Config{
		AgentAPISocket: getEnvOrDefault("PIXCONF_AGENT_API_SOCKET", "/var/run/pixconf.sock"),
		AgentID:        getEnvOrDefaultFunc("PIXCONF_AGENT_ID", getHostname),
		Server:         getEnvOrDefaultFunc("PIXCONF_SERVER", getServer),
	}
}
