package mqtthandler

import (
	"github.com/pixconf/pixconf/internal/agentmeta"
	"github.com/pixconf/pixconf/pkg/mqttmsg"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *Hook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	var validSignature bool

	if pk.Properties.User != nil {
		for _, row := range pk.Properties.User {
			if row.Key == mqttmsg.HeaderPayloadSignature {
				validSignature = true
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

	if pk.Properties.CorrelationData == nil {
		h.log.Warn(
			"missing correlation data",
			"client", cl.ID,
			"topic", pk.TopicName,
		)
		return
	}

	topics := agentmeta.GetTopics(cl.ID)

	topicRespone := agentmeta.GetResponseTopic(cl.ID, string(pk.Properties.CorrelationData))

	switch pk.TopicName {
	case topics.Health:
		h.log.Debug(
			"health",
			"client", cl.ID,
			"topic", pk.TopicName,
			"payload", string(pk.Payload),
		)

	case topicRespone:
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
