package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/pixconf/pixconf/pkg/mqttmsg"
	"github.com/pixconf/pixconf/pkg/server/proto"
)

func (app *Agent) mqttTopicCommandHandler(pb *paho.Publish) {
	fmt.Printf(
		"received message with topic: %s | payload: %s\n",
		pb.Topic,
		string(pb.Payload),
	)

	if pb.Properties == nil {
		app.log.Error("received command message without properties")
		return
	}

	if pb.Properties.ContentType != mqttmsg.ContentTypeJSON {
		app.log.Error("received command message with invalid content type")
		return
	}

	var request proto.AgentRPCRequest

	startTime := time.Now()

	if err := json.Unmarshal(pb.Payload, &request); err != nil {
		app.log.Error("failed to unmarshal request", "error", err)
		return
	}

	// response only if the message has a response topic
	if pb.Properties != nil && len(pb.Properties.ResponseTopic) > 0 {
		response := &proto.AgentRPCResponse{
			RequestID:     request.RequestID,
			Request:       &request,
			ExecutionTime: time.Since(startTime).Seconds(),
		}

		responsePayload, err := json.Marshal(response)
		if err != nil {
			app.log.Error("failed to marshal response", "error", err)
			return
		}

		publish := &autopaho.QueuePublish{
			Publish: &paho.Publish{
				Topic:   pb.Properties.ResponseTopic,
				Payload: responsePayload,

				Properties: &paho.PublishProperties{
					ContentType:     mqttmsg.ContentTypeJSON,
					CorrelationData: pb.Properties.CorrelationData,
				},
			},
		}

		if err := app.mqttConn.PublishViaQueue(context.Background(), publish); err != nil {
			log.Println("failed to publish response message:", err)
		}
	}
}
