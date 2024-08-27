package mqttauth

import (
	"bytes"
	"log/slog"

	mqtt "github.com/mochi-mqtt/server/v2"
)

type Hook struct {
	mqtt.HookBase
	log    *slog.Logger
	config *HookOptions
}

type HookOptions struct {
	Server *mqtt.Server
}

func NewHook(logger *slog.Logger) *Hook {
	return &Hook{
		log: logger,
	}
}

func (h *Hook) ID() string {
	return "agent-auth"
}

func (h *Hook) Provides(b byte) bool {
	provides := []byte{
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
	}

	return bytes.Contains(provides, []byte{b})
}

func (h *Hook) Init(config any) error {
	h.log.Debug("hook", "Init", "config", config)

	if _, ok := config.(*HookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	h.config = config.(*HookOptions)

	return nil
}
