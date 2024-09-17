package xkit

import (
	"fmt"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/rs/xid"
)

type MQTTPublishRequest struct {
	Topic      string
	Payload    []byte
	Properties packets.Properties
	Expiry     time.Duration
}

// hack: https://github.com/mochi-mqtt/server/issues/428
func MQTTPublish(server *mqtt.Server, request *MQTTPublishRequest) error {
	clientID := fmt.Sprintf("inline-%s", xid.New().String())
	client := server.NewClient(nil, "local", clientID, true)

	packet := packets.Packet{
		FixedHeader: packets.FixedHeader{
			Type: packets.Publish,
		},
		TopicName:  request.Topic,
		Payload:    request.Payload,
		PacketID:   0,
		Properties: request.Properties,
		Created:    time.Now().Unix(),
	}

	if request.Expiry > 0 {
		packet.Expiry = time.Now().Add(request.Expiry).Unix()
	}

	return server.InjectPacket(client, packet)
}
