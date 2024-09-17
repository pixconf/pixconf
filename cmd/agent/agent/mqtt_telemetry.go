package agent

import (
	"context"
	"encoding/json"
	"time"

	"github.com/eclipse/paho.golang/paho"
	"github.com/google/uuid"
	"github.com/pixconf/pixconf/internal/agentmeta"
	"github.com/pixconf/pixconf/pkg/mqttmsg"
	"github.com/pixconf/pixconf/pkg/xkit"
)

func (app *Agent) mqttSendHealthTelemetry(ctx context.Context) error {
	time.Sleep(xkit.RandomTime(30))

	if app.mqttConn == nil {
		return nil
	}

	telemetry := mqttmsg.AgentTelemetry{
		AgentID:     app.config.AgentID,
		StartedTime: app.startedTime.Unix(),
	}

	payload, err := json.Marshal(telemetry)
	if err != nil {
		return err
	}

	requestID := uuid.New().String()

	topics := agentmeta.GetTopics(app.config.AgentID)

	publish := &paho.Publish{
		Topic:   topics.Health,
		Payload: payload,
		Properties: &paho.PublishProperties{
			ContentType:     mqttmsg.ContentTypeJSON,
			CorrelationData: []byte(requestID),
		},
	}

	if _, err := app.mqttConn.Publish(ctx, publish); err != nil {
		return err
	}

	return nil
}
