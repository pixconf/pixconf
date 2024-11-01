package mqttauth

import (
	"log/slog"
	"testing"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/stretchr/testify/require"
)

func TestHookOnACLCheck(t *testing.T) {
	h := NewHook(slog.New(slog.NewTextHandler(nil, nil)))

	testCases := []struct {
		name     string
		topic    string
		write    bool
		expected bool
	}{
		{
			name:     "allow read from commands topic",
			topic:    "pixconf/agent/username/commands",
			write:    false,
			expected: true,
		},
		{
			name:     "deny write to commands topic",
			topic:    "pixconf/agent/username/commands",
			write:    true,
			expected: false,
		},
		{
			name:     "allow write to health topic",
			topic:    "pixconf/agent/username/health",
			write:    true,
			expected: true,
		},
		{
			name:     "deny read from health topic",
			topic:    "pixconf/agent/username/health",
			write:    false,
			expected: false,
		},
		{
			name:     "deny read from response topic",
			topic:    "pixconf/agent/username/response/bd07249d-c967-4f42-87c3-e2dccc03a3ff",
			write:    false,
			expected: false,
		},
		{
			name:     "allow write to correct response topic",
			topic:    "pixconf/agent/username/response/bd07249d-c967-4f42-87c3-e2dccc03a3ff",
			write:    true,
			expected: true,
		},
		{
			name:     "deny write to incorrect response topic",
			topic:    "pixconf/agent/username/response/incorrect",
			write:    true,
			expected: false,
		},
		{
			name:     "deny write to incorrect response topic with pattern",
			topic:    "pixconf/agent/username/response/????????-????-????-????-????????????",
			write:    true,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := &mqtt.Client{}
			client.Properties.Username = []byte("username")

			require.Equal(t, tc.expected, h.OnACLCheck(client, tc.topic, tc.write))
		})
	}
}
