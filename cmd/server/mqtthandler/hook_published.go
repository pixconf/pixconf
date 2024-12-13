package mqtthandler

import (
	"github.com/pixconf/pixconf/internal/agentmeta"
	"github.com/pixconf/pixconf/pkg/mqttmsg"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *Hook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	var validSignature bool
	var requestID string

	if pk.Properties.User != nil {
		for _, row := range pk.Properties.User {
			if row.Key == mqttmsg.HeaderPayloadSignature {
				validSignature = true
			}

			if row.Key == mqttmsg.HeaderRequestID {
				requestID = row.Val
			}
		}
	}

	if !validSignature {
		h.log.Warn(
			"invalid signature",
			"client", cl.ID,
			"topic", pk.TopicName,
		)
		return
	}

	topics := agentmeta.GetTopics(cl.ID)

	switch pk.TopicName {
	case topics.Health:
		h.log.Debug(
			"health",
			"client", cl.ID,
			"topic", pk.TopicName,
			"payload", string(pk.Payload),
		)

	case agentmeta.GetResponseTopic(cl.ID, requestID):
		h.log.Debug(
			"response",
			"client", cl.ID,
			"topic", pk.TopicName,
			"payload", string(pk.Payload),
		)

	default:
		h.log.Warn(
			"unknown topic",
			"client", cl.ID,
			"topic", pk.TopicName,
		)
	}
}
