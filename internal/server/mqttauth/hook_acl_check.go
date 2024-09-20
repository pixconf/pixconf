package mqttauth

import (
	"fmt"
	"strings"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/pixconf/pixconf/internal/buildinfo"
	"github.com/tidwall/match"
)

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
