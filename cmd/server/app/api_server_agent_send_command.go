package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/pixconf/pixconf/cmd/server/mqttkit"
	"github.com/pixconf/pixconf/internal/agentmeta"
	"github.com/pixconf/pixconf/internal/apitool"
	"github.com/pixconf/pixconf/pkg/mqttmsg"
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

	requestID := xkit.GetUUID(apitool.GetRequestID(c))

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

	agentTopics := agentmeta.GetTopics(content.Agent)

	responseTopic := agentmeta.GetResponseTopic(content.Agent, request.RequestID)

	mqttRequest := &mqttkit.MQTTPublishRequest{
		Topic:   agentTopics.Commands,
		Payload: requestPayload,
		Properties: packets.Properties{
			CorrelationData: xkit.GetUUIDBytes(requestID),
			ContentType:     mqttmsg.ContentTypeJSON,
			ResponseTopic:   responseTopic,
		},
	}

	var wait sync.WaitGroup
	wait.Add(1)

	if err := mqttkit.MQTTPublish(app.mqtt, mqttRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var response proto.AgentRPCResponse

	callbackFn := func(_ *mqtt.Client, _ packets.Subscription, pk packets.Packet) {
		// BUG: panic on duplicate response
		defer wait.Done()

		if pk.Properties.ContentType != mqttmsg.ContentTypeJSON {
			return
		}

		if err := json.Unmarshal(pk.Payload, &response); err != nil {
			fmt.Printf("Failed to unmarshal response: %s\n", err)
			return
		}
	}

	subscriptionID := int(time.Now().Unix())

	app.mqtt.Subscribe(responseTopic, subscriptionID, callbackFn)

	if xkit.WaitTimeout(&wait, 10*time.Second) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "command timeout"})
		return
	}

	app.mqtt.Unsubscribe(responseTopic, subscriptionID)

	c.JSON(http.StatusOK, gin.H{"response": response})
}
