package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eclipse/paho.golang/paho"
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

	publish := &paho.Publish{
		Topic:   fmt.Sprintf("pixconf/agent/%s/health", app.mqttClientID),
		QoS:     0x0,
		Payload: payload,
	}

	if _, err := app.mqttConn.Publish(ctx, publish); err != nil {
		return err
	}

	return nil
}
