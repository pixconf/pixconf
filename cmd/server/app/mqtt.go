package app

import (
	"math"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/pixconf/pixconf/cmd/server/mqttauth"
	"github.com/pixconf/pixconf/internal/buildinfo"
)

func (app *App) initMQTT() error {
	caps := mqtt.NewDefaultServerCapabilities()
	caps.MaximumClients = math.MaxUint16 / 2       // Maximum number of clients (emm, if we reach this number, we are in trouble) (32767)
	caps.MaximumSessionExpiryInterval = 14 * 86400 // Maximum number of seconds to keep disconnected sessions (14 days, default - 136 years)
	caps.MinimumProtocolVersion = 5                // Support only MQTT 5.0
	caps.MaximumQos = 0x0                          // Support QoS 0, 1, 2
	caps.RetainAvailable = 0x0                     // Support retain messages
	caps.SharedSubAvailable = 0x0                  // Support shared subscriptions

	caps.Compatibilities.RestoreSysInfoOnRestart = false

	caps.Compatibilities.ObscureNotAuthorized = false      // for paho.mqtt.golang (maybe not needed)
	caps.Compatibilities.PassiveClientDisconnect = true    // for paho.mqtt.golang
	caps.Compatibilities.NoInheritedPropertiesOnAck = true // for paho.mqtt.golang

	options := &mqtt.Options{
		Logger:       app.logger.With("service", "mqtt"),
		Capabilities: caps,
		InlineClient: true,
	}

	app.mqtt = mqtt.New(options)
	app.mqtt.Info.Version = buildinfo.Version

	mqttAuthHook := mqttauth.NewHook(app.logger.With("service", "mqtt-auth"))

	if err := app.mqtt.AddHook(mqttAuthHook, &mqttauth.HookOptions{
		Server: app.mqtt,
	}); err != nil {
		return err
	}

	tcp := listeners.NewTCP(listeners.Config{
		ID:      "mqtt-agent",
		Address: ":1883",
	})

	if err := app.mqtt.AddListener(tcp); err != nil {
		return err
	}

	return nil
}

func (app *App) ListenAndServeMQTT() error {
	return app.mqtt.Serve()
}
