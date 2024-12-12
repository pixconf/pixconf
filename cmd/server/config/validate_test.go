package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		conf    Config
		wantErr bool
	}{
		{
			name: "valid config",
			conf: Config{
				MQTT: MQTTConfig{Listen: []string{"mqtt://localhost:1883"}},
			},
			wantErr: false,
		},
		{
			name: "missing listen",
			conf: Config{
				MQTT: MQTTConfig{Listen: nil},
			},
			wantErr: true,
		},
		{
			name: "invalid scheme",
			conf: Config{
				MQTT: MQTTConfig{Listen: []string{"http://localhost:1883"}},
			},
			wantErr: true,
		},
		{
			name: "invalid URL",
			conf: Config{
				MQTT: MQTTConfig{Listen: []string{"::invalid-url::"}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.conf); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
