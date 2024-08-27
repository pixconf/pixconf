package mqttauth

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *Hook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	if !h.limiter.Allow() {
		return packets.ErrServerBusy
	}

	return nil
}
