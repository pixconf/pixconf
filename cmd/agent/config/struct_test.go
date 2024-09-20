package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name           string
		envVars        map[string]string
		expectedConfig Config
	}{
		{
			name:    "DefaultValues",
			envVars: map[string]string{},
			expectedConfig: Config{
				AgentAPISocket: "/var/run/pixconf.sock",
				AgentID:        getHostname(),
				AuthKeyPath:    "/etc/pixconf/agent_auth.key",
				Server:         getServer(),
			},
		},
		{
			name: "CustomAgentAPISocket",
			envVars: map[string]string{
				"PIXCONF_AGENT_API_SOCKET": "/custom/socket",
			},
			expectedConfig: Config{
				AgentAPISocket: "/custom/socket",
				AgentID:        getHostname(),
				AuthKeyPath:    "/etc/pixconf/agent_auth.key",
				Server:         getServer(),
			},
		},
		{
			name: "CustomAgentID",
			envVars: map[string]string{
				"PIXCONF_AGENT_ID": "custom_id",
			},
			expectedConfig: Config{
				AgentAPISocket: "/var/run/pixconf.sock",
				AgentID:        "custom_id",
				AuthKeyPath:    "/etc/pixconf/agent_auth.key",
				Server:         getServer(),
			},
		},
		{
			name: "CustomServer",
			envVars: map[string]string{
				"PIXCONF_SERVER": "custom_server",
			},
			expectedConfig: Config{
				AgentAPISocket: "/var/run/pixconf.sock",
				AgentID:        getHostname(),
				AuthKeyPath:    "/etc/pixconf/agent_auth.key",
				Server:         "custom_server",
			},
		},
		{
			name: "AllCustomValues",
			envVars: map[string]string{
				"PIXCONF_AGENT_API_SOCKET": "/custom/socket",
				"PIXCONF_AGENT_ID":         "custom_id",
				"PIXCONF_AUTH_KEY_PATH":    "/custom/path",
				"PIXCONF_SERVER":           "custom_server",
			},
			expectedConfig: Config{
				AgentAPISocket: "/custom/socket",
				AgentID:        "custom_id",
				AuthKeyPath:    "/custom/path",
				Server:         "custom_server",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}

			// Create new config
			config := newConfig()

			// Assert the config matches the expected config
			assert.Equal(t, tt.expectedConfig, *config)
		})
	}
}

func TestIsZeroValue(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  bool
	}{
		{"ZeroInt", 0, true},
		{"NonZeroInt", 1, false},
		{"ZeroString", "", true},
		{"NonZeroString", "test", false},
		{"ZeroBool", false, true},
		{"NonZeroBool", true, false},
		{"ZeroStruct", Config{}, true},
		{"NonZeroStruct", Config{AgentID: "123"}, false},
		{"ZeroPointer", (*Config)(nil), true},
		{"NonZeroPointer", &Config{AgentID: "123"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := reflect.ValueOf(tt.value)
			got := isZeroValue(v)

			assert.Equal(t, tt.want, got)
		})

	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name           string
		defaultConfig  Config
		customConfig   Config
		expectedConfig Config
	}{
		{
			name:           "MergeWithEmptyCustomConfig",
			defaultConfig:  Config{AgentAPISocket: "default_socket", AgentID: "default_id", Server: "default_server"},
			customConfig:   Config{},
			expectedConfig: Config{AgentAPISocket: "default_socket", AgentID: "default_id", Server: "default_server"},
		},
		{
			name:           "MergeWithPartialCustomConfig",
			defaultConfig:  Config{AgentAPISocket: "default_socket", AgentID: "default_id", Server: "default_server"},
			customConfig:   Config{AgentID: "custom_id"},
			expectedConfig: Config{AgentAPISocket: "default_socket", AgentID: "custom_id", Server: "default_server"},
		},
		{
			name:           "MergeWithFullCustomConfig",
			defaultConfig:  Config{AgentAPISocket: "default_socket", AgentID: "default_id", Server: "default_server"},
			customConfig:   Config{AgentAPISocket: "custom_socket", AgentID: "custom_id", Server: "custom_server"},
			expectedConfig: Config{AgentAPISocket: "custom_socket", AgentID: "custom_id", Server: "custom_server"},
		},
		{
			name:           "MergeWithNilCustomConfig",
			defaultConfig:  Config{AgentAPISocket: "default_socket", AgentID: "default_id", Server: "default_server"},
			customConfig:   Config{},
			expectedConfig: Config{AgentAPISocket: "default_socket", AgentID: "default_id", Server: "default_server"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultConfig := tt.defaultConfig
			customConfig := tt.customConfig

			defaultConfig.merge(&customConfig)

			assert.Equal(t, tt.expectedConfig, defaultConfig)
		})
	}
}
