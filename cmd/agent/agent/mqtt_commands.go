package agent

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/pixconf/pixconf/pkg/mqttmsg"
	"github.com/pixconf/pixconf/pkg/server/proto"
)

func (app *Agent) mqttTopicCommandHandler(pb *paho.Publish) {
	app.log.Debug("received command message", "topic", pb.Topic, "payload", string(pb.Payload), "properties", pb.Properties)

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

		publishProperties := &paho.PublishProperties{
			ContentType:     mqttmsg.ContentTypeJSON,
			CorrelationData: pb.Properties.CorrelationData,
		}

		publishProperties.User.Add(mqttmsg.HeaderRequestID, request.RequestID)

		publish := &autopaho.QueuePublish{
			Publish: &paho.Publish{
				Topic:   pb.Properties.ResponseTopic,
				Payload: responsePayload,

				Properties: publishProperties,
			},
		}

		if err := app.mqttConn.PublishViaQueue(context.Background(), publish); err != nil {
			log.Println("failed to publish response message:", err)
		}
	}
}
