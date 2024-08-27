package mqttauth

import (
	"log/slog"
	"testing"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/stretchr/testify/require"
)

func TestHookOnConnectAuthenticate(t *testing.T) {
	h := NewHook(slog.New(slog.NewTextHandler(nil, nil)))

	client := new(mqtt.Client)
	client.Net.Remote = "192.0.2.11"

	require.True(t, h.OnConnectAuthenticate(client, packets.Packet{}))
}
