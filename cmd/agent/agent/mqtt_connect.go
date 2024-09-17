package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/pixconf/pixconf/pkg/agent/agent2server"
)

func (app *Agent) mqttConnect(ctx context.Context) error {
	cliCfg := app.mqttConnectConfig()

	serverClient, err := agent2server.NewClient(agent2server.Options{
		ServerEndpoint: app.config.Server,
	})
	if err != nil {
		return err
	}

	agentAutoConfiguration, err := serverClient.GetAgentAutoConfiguration(ctx)
	if err != nil {
		return err
	}

	for _, serverURL := range agentAutoConfiguration.MQTTEndpoints {
		parsedURL, err := url.Parse(serverURL)
		if err != nil {
			continue
		}

		cliCfg.ServerUrls = append(cliCfg.ServerUrls, parsedURL)
	}

	app.mqttConn, err = autopaho.NewConnection(ctx, cliCfg)
	if err != nil {
		return err
	}

	if err := app.mqttConn.AwaitConnection(ctx); err != nil {
		return err
	}

	<-app.mqttConn.Done()

	return nil
}

func (app *Agent) mqttConnectConfig() autopaho.ClientConfig {
	topics := getMQTTTopics(app.config.AgentID)

	config := autopaho.ClientConfig{
		KeepAlive:                     30,
		CleanStartOnInitialConnection: true,
		SessionExpiryInterval:         60, // session remains live 60 seconds after disconnect
		Queue:                         app.mqttQueue,
		ConnectUsername:               app.config.AgentID,
		ConnectPassword:               []byte("my.great.jwt.token.here.123"), // TODO: replace with actual JWT token
	}

	config.OnConnectionUp = func(cm *autopaho.ConnectionManager, _ *paho.Connack) {
		app.log.Info("connected to MQTT broker")

		if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
			Subscriptions: []paho.SubscribeOptions{
				{Topic: topics.Commands, QoS: 0x0},
			},
		}); err != nil {
			app.log.Warn("error whilst subscribing:", "error", err)
		}
	}

	config.OnConnectError = func(err error) {
		app.log.Warn("error whilst attempting connection:", "error", err)
	}

	config.ClientConfig = paho.ClientConfig{
		ClientID: app.config.AgentID,
	}

	config.ClientConfig.OnPublishReceived = []func(paho.PublishReceived) (bool, error){
		func(pr paho.PublishReceived) (bool, error) {
			switch pr.Packet.Topic {
			case topics.Commands:
				app.mqttTopicCommandHandler(pr.Packet)

			default:
				app.log.Warn("received message on unknown topic", "topic", pr.Packet.Topic)
				return false, fmt.Errorf("unknown topic: %s", pr.Packet.Topic)
			}

			return false, nil
		},
	}

	config.ClientConfig.OnServerDisconnect = func(d *paho.Disconnect) {
		if d.Properties != nil {
			app.log.Warn("disconnected from MQTT server", "reason", d.Properties.ReasonString)
		} else {
			app.log.Warn("disconnected from MQTT server", "reason-code", d.ReasonCode)
		}
	}

	return config
}
