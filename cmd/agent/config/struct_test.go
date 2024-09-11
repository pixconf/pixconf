package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
