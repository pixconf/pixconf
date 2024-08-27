package mqttauth

import (
	"testing"

	"github.com/mochi-mqtt/server/v2/packets"
)

func TestOnConnect(t *testing.T) {
	h := NewHook(nil)

	maxRate := 10

	for i := 0; i < 100; i++ {
		if err := h.OnConnect(nil, packets.Packet{}); err != nil {
			if i < maxRate {
				t.Errorf("expected no error, got %v", err)
			}

			if err != packets.ErrServerBusy {
				t.Errorf("expected error %v, got %v", packets.ErrServerBusy, err)
			}
		}
	}
}
