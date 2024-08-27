package mqttauth

import (
	"testing"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/stretchr/testify/require"
)

func TestHookID(t *testing.T) {
	h := new(Hook)
	require.Equal(t, "agent-auth", h.ID())
}

func TestHookProvides(t *testing.T) {
	h := new(Hook)

	require.True(t, h.Provides(mqtt.OnACLCheck))
	require.True(t, h.Provides(mqtt.OnConnectAuthenticate))
	require.False(t, h.Provides(mqtt.OnPublished))
}
