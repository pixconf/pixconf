package mqttauth

import (
	"bytes"
	"fmt"
	"log/slog"
	"strings"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/pixconf/pixconf/internal/buildinfo"
	"github.com/tidwall/match"
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

func (h *Hook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	username := string(cl.Properties.Username)

	var (
		allowMessage   bool
		continueChecks = true
	)

	// only for debug purposes under development
	if buildinfo.Version == "dev" && strings.HasPrefix(topic, "$SYS/") {
		allowMessage = true
		continueChecks = false
	}

	// deny topics with pattern characters - this case is weird, so we cick the client
	for _, row := range []string{"#", "?", "+"} {
		if continueChecks && strings.Contains(topic, row) {

			if h.config != nil && h.config.Server != nil {
				// https://developers.cloudflare.com/pub-sub/platform/mqtt-compatibility/
				h.config.Server.DisconnectClient(cl, packets.ErrTopicNameInvalid)
			}

			continueChecks = false
			break
		}
	}

	if continueChecks {
		// map: topic -> write
		accessTopics := map[string]bool{
			fmt.Sprintf("pixconf/agent/%s/commands", username): false,
			fmt.Sprintf("pixconf/agent/%s/health", username):   true,

			fmt.Sprintf("pixconf/agent/%s/response/????????-????-????-????-????????????", username): true,
		}

		for topicPattern, w := range accessTopics {
			if topicPattern == topic {
				allowMessage = w == write
				break
			}

			if match.IsPattern(topicPattern) && match.Match(topic, topicPattern) {
				allowMessage = w == write
				break
			}
		}
	}

	requestType := "subscribe"
	if write {
		requestType = "publish"
	}

	h.log.Debug(
		fmt.Sprintf("%v", allowMessage),
		"hook", "OnACLCheck",
		"username", username,
		"topic", topic,
		"request-type", requestType,
	)

	return allowMessage
}
