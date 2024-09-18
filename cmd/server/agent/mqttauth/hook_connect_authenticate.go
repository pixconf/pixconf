package mqttauth

import (
	"strings"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/pixconf/pixconf/internal/agent/authkey"
)

func (h *Hook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	if pk.Connect.Username == nil || pk.Connect.Password == nil {
		if h.config != nil {
			h.config.Server.DisconnectClient(cl, packets.ErrBadUsernameOrPassword)
		}
		return false
	}

	username := string(pk.Connect.Username)
	password := string(pk.Connect.Password)

	// username must be lowercase
	if strings.ToLower(username) != username {
		if h.config != nil {
			h.config.Server.DisconnectClient(cl, packets.ErrClientIdentifierNotValid)
		}
		return false
	}

	// username must be equal to client_id
	if cl.ID != username {
		if h.config != nil {
			h.config.Server.DisconnectClient(cl, packets.ErrClientIdentifierNotValid)
		}
		return false
	}

	// username must be less than 255 characters. like domain name
	if len(username) > 255 {
		if h.config != nil {
			h.config.Server.DisconnectClient(cl, packets.ErrProtocolViolationUsernameTooLong)
		}
		return false
	}

	_, payload, err := authkey.Unmarshal(password)
	if err != nil {
		if h.config != nil {
			h.config.Server.DisconnectClient(cl, packets.ErrBadUsernameOrPassword)
		}
		return false
	}

	if payload.Issuer != username {
		if h.config != nil {
			h.config.Server.DisconnectClient(cl, packets.ErrClientIdentifierNotValid)
		}
		return false
	}

	return false
}
