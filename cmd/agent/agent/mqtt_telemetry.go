package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eclipse/paho.golang/paho"
	"github.com/google/uuid"
	"github.com/pixconf/pixconf/pkg/mqttmsg"
	"github.com/pixconf/pixconf/pkg/xkit"
)

func (app *Agent) mqttSendHealthTelemetry(ctx context.Context) error {
	time.Sleep(xkit.RandomTime(30))

	if app.mqttConn == nil {
		return nil
	}

	telemetry := mqttmsg.AgentTelemetry{
		AgentID:     app.mqttClientID,
		StartedTime: app.startedTime.Unix(),
	}

	payload, err := json.Marshal(telemetry)
	if err != nil {
		return err
	}

	requestID := uuid.New().String()

	publish := &paho.Publish{
		Topic:   fmt.Sprintf("pixconf/agent/%s/health", app.mqttClientID),
		Payload: payload,
		Properties: &paho.PublishProperties{
			ContentType:     "application/json",
			CorrelationData: []byte(requestID),
		},
	}

	if _, err := app.mqttConn.Publish(ctx, publish); err != nil {
		return err
	}

	return nil
}
