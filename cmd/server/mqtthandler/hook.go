package mqtthandler

import (
	"bytes"
	"fmt"
	"log/slog"

	mqtt "github.com/mochi-mqtt/server/v2"
)

type Hook struct {
	mqtt.HookBase
	log *slog.Logger
}

func NewHook(logger *slog.Logger) *Hook {
	return &Hook{
		log: logger,
	}
}

func (h *Hook) ID() string {
	return "agent-handler"
}

func (h *Hook) Provides(b byte) bool {
	provides := []byte{
		mqtt.OnPublished,
	}

	return bytes.Contains(provides, []byte{b})
}

func (h *Hook) Init(config any) error {
	h.log.Debug(
		fmt.Sprintf("%v", config),
		"hook", "init",
	)

	return nil
}
