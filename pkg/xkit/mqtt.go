package xkit

import (
	"fmt"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/rs/xid"
)

type MQTTPublishRequest struct {
	Topic      string
	Payload    []byte
	Properties packets.Properties
}

// hack: https://github.com/mochi-mqtt/server/issues/428
func MQTTPublish(server *mqtt.Server, request *MQTTPublishRequest) error {
	clientID := fmt.Sprintf("inline-%s", xid.New().String())
	client := server.NewClient(nil, "local", clientID, true)

	return server.InjectPacket(client, packets.Packet{
		FixedHeader: packets.FixedHeader{
			Type: packets.Publish,
		},
		TopicName:  request.Topic,
		Payload:    request.Payload,
		PacketID:   0,
		Properties: request.Properties,
	})
}
