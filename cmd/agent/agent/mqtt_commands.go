package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/eclipse/paho.golang/paho"
	"github.com/pixconf/pixconf/pkg/server/proto"
)

func (app *Agent) mqttTopicCommandHandler(pb *paho.Publish) {
	fmt.Printf(
		"received message with topic: %s | payload: %s\n",
		pb.Topic,
		string(pb.Payload),
	)

	var request proto.AgentRPCRequest

	if err := json.Unmarshal(pb.Payload, &request); err != nil {
		app.log.Error("failed to unmarshal request", "error", err)
		return
	}

	// response only if the message has a response topic
	if pb.Properties != nil && len(pb.Properties.ResponseTopic) > 0 {
		publish := &paho.Publish{
			Topic: pb.Properties.ResponseTopic,
			Properties: &paho.PublishProperties{
				CorrelationData: pb.Properties.CorrelationData,
			},
		}

		publish.Payload = pb.Payload

		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if _, err := app.mqttConn.Publish(ctxTimeout, publish); err != nil {
			log.Println("failed to publish response message:", err)
		}
	}
}