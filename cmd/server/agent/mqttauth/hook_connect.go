package mqttauth

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *Hook) OnConnect(_ *mqtt.Client, _ packets.Packet) error {
	if !h.limiter.Allow() {
		return packets.ErrServerBusy
	}

	return nil
}
