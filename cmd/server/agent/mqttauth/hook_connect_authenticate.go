package mqttauth

import (
	"fmt"
	"strings"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *Hook) OnConnectAuthenticate(cl *mqtt.Client, _ packets.Packet) bool {
	var allowConnect = true

	username := string(cl.Properties.Username)

	// username must be lowercase
	if strings.ToLower(username) != username {
		allowConnect = false
	}

	// username must be equal to client_id
	if cl.ID != username {
		allowConnect = false
	}

	// username must be less than 255 characters. like domain name
	if len(username) > 255 {
		allowConnect = false
	}

	h.log.Debug(
		fmt.Sprintf("%v", allowConnect),
		"hook", "OnConnectAuthenticate",
		"username", username,
		"remote_addr", cl.Net.Remote,
	)

	return allowConnect
}
