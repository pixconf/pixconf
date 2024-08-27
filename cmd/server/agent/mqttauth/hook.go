package mqttauth

import (
	"bytes"
	"fmt"
	"log/slog"

	mqtt "github.com/mochi-mqtt/server/v2"
	"golang.org/x/time/rate"
)

type Hook struct {
	mqtt.HookBase
	log    *slog.Logger
	config *HookOptions

	limiter *rate.Limiter
}

type HookOptions struct {
	Server *mqtt.Server
}

func NewHook(logger *slog.Logger) *Hook {
	// TODO: make rate limit configurable or dynamic
	// TODO: make rate limit based on client ID or IP
	limiter := rate.NewLimiter(rate.Limit(10), 10) // 10 requests per second

	return &Hook{
		log:     logger,
		limiter: limiter,
	}
}

func (h *Hook) ID() string {
	return "agent-auth"
}

func (h *Hook) Provides(b byte) bool {
	provides := []byte{
		mqtt.OnConnect,
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
	}

	return bytes.Contains(provides, []byte{b})
}

func (h *Hook) Init(config any) error {
	h.log.Debug(
		fmt.Sprintf("%v", config),
		"hook", "init",
	)

	if _, ok := config.(*HookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	h.config = config.(*HookOptions)

	return nil
}
