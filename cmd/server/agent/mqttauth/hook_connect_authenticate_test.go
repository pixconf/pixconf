package mqttauth

import (
	"log/slog"
	"strings"
	"testing"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/stretchr/testify/require"
)

func TestHookOnConnectAuthenticate(t *testing.T) {
	h := NewHook(slog.New(slog.NewTextHandler(nil, nil)))

	const (
		validUser  = "4f8ab175b616"
		validToken = "eyJhbGciOiJlZDI1NTE5IiwicGsiOiJQbzdVS1lzaUJJb1d6M2djaGJocTdtQVI5SDhhQjlWandIcHFlU2RObFhjIn0.eyJpc3MiOiI0ZjhhYjE3NWI2MTYiLCJqdGkiOiI4ZDE0MGE2Ni0wNDgyLTQ1YjktYjVmOC1kZjJhNDY5MzNhOTEiLCJpYXQiOjE3MjY2NDIwNDksInZlciI6ImRldiJ9.QfzxaGEQSBrTzm3riEwHpnUFQTbfby68fbNg7pVIce2C9Gtl8UADJkKoe5AXlxy6WIDdXypNGxyaWd_6uPGPBw"
	)

	t.Run("empty_username", func(t *testing.T) {
		client := new(mqtt.Client)

		require.False(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Password: []byte("password"),
			},
		}))
	})

	t.Run("empty_password", func(t *testing.T) {
		client := new(mqtt.Client)

		require.False(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Username: []byte("username"),
			},
		}))
	})

	t.Run("invalid_username", func(t *testing.T) {
		client := new(mqtt.Client)

		require.False(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Username: []byte("Username"),
				Password: []byte("password"),
			},
		}))
	})

	t.Run("invalid_client_id", func(t *testing.T) {
		client := new(mqtt.Client)

		require.False(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Username: []byte("username"),
				Password: []byte("password"),
			},
		}))
	})

	t.Run("username_too_long", func(t *testing.T) {
		client := new(mqtt.Client)
		client.ID = strings.Repeat("a", 256)

		require.False(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Username: []byte(client.ID),
				Password: []byte("password"),
			},
		}))
	})

	t.Run("invalid_token", func(t *testing.T) {
		client := new(mqtt.Client)
		client.ID = validUser

		require.False(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Username: []byte(validUser),
				Password: []byte("invalid"),
			},
		}))
	})

	t.Run("valid", func(t *testing.T) {
		client := new(mqtt.Client)
		client.ID = validUser

		require.True(t, h.OnConnectAuthenticate(client, packets.Packet{
			Connect: packets.ConnectParams{
				Username: []byte(validUser),
				Password: []byte(validToken),
			},
		}))
	})

}
