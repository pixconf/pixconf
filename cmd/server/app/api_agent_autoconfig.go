package app

import (
	"fmt"
	"net/http"
	"os"
	"slices"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/pixconf/pixconf/pkg/server/proto"
)

func (app *App) apiAgentAutoConfiguration(ctx *gin.Context) {
	config := proto.AgentAutoConfigurationResponse{
		ClientIPAddress: ctx.ClientIP(),
	}

	if app.config.MQTT.Endpoints != nil {
		config.MQTTEndpoints = app.config.MQTT.Endpoints

	} else {
		// TODO: make this fallback more robust

		currentURL := location.Get(ctx)

		endpointNames := []string{
			currentURL.Hostname(),
		}

		if serverName, err := os.Hostname(); err == nil {
			endpointNames = append(endpointNames, serverName)
		}

		slices.Sort(endpointNames)
		endpointNames = slices.Compact(endpointNames)

		switch currentURL.Scheme {
		case "http":
			for _, name := range endpointNames {
				config.MQTTEndpoints = append(config.MQTTEndpoints, fmt.Sprintf("tcp://%s:1883", name))
			}

		case "https":
			for _, name := range endpointNames {
				config.MQTTEndpoints = append(config.MQTTEndpoints, fmt.Sprintf("tls://%s:8883", name))
			}
		}

	}

	ctx.JSON(http.StatusOK, config)
}
