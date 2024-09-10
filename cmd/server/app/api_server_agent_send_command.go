package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/pixconf/pixconf/pkg/server/proto"
	"github.com/pixconf/pixconf/pkg/xkit"
)

type apiServerAgentSendCommandRequest struct {
	Agent   string            `json:"agent" binding:"required"`
	Command string            `json:"command" binding:"required"`
	Args    []string          `json:"args"`
	KWArgs  map[string]string `json:"kwargs"`
}

func (app *App) apiServerAgentSendCommand(c *gin.Context) {
	var content apiServerAgentSendCommandRequest

	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := app.mqtt.Clients.Get(content.Agent); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("agent not found: %s", content.Agent)})
		return
	}

	requestID := xkit.GetUUID(c.GetHeader("X-Request-ID"))

	request := &proto.AgentRPCRequest{
		RequestID: requestID,
		Command:   content.Command,
		Args:      content.Args,
		KWArgs:    content.KWArgs,
	}

	requestPayload, err := json.Marshal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mqttRequest := &xkit.MQTTPublishRequest{
		Topic:   fmt.Sprintf("pixconf/agent/%s/commands", content.Agent),
		Payload: requestPayload,
		Properties: packets.Properties{
			CorrelationData: []byte(request.RequestID),
			ContentType:     "application/json",
			ResponseTopic:   fmt.Sprintf("pixconf/agent/%s/response/%s", content.Agent, request.RequestID),
		},
	}

	if err := xkit.MQTTPublish(app.mqtt, mqttRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
