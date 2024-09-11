package agent

import (
	"context"
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

	router := paho.NewStandardRouter()
	router.RegisterHandler(topics.Commands, app.mqttTopicCommandHandler)

	config := autopaho.ClientConfig{
		KeepAlive:                     30,
		CleanStartOnInitialConnection: true,
		SessionExpiryInterval:         60, // session remains live 60 seconds after disconnect
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
		OnPublishReceived: []func(paho.PublishReceived) (bool, error){
			func(pr paho.PublishReceived) (bool, error) {
				router.Route(pr.Packet.Packet())
				return true, nil // we assume that the router handles all messages
			},
		},
	}

	return config
}
