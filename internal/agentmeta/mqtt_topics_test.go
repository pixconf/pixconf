package agentmeta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopics(t *testing.T) {
	tests := []struct {
		agentID        string
		expectedTopics Topics
	}{
		{
			agentID: "agent123",
			expectedTopics: Topics{
				Commands: "pixconf/agent/agent123/commands",
				Health:   "pixconf/agent/agent123/health",
			},
		},
		{
			agentID: "agent456",
			expectedTopics: Topics{
				Commands: "pixconf/agent/agent456/commands",
				Health:   "pixconf/agent/agent456/health",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.agentID, func(t *testing.T) {
			got := GetTopics(tt.agentID)

			assert.Equal(t, tt.expectedTopics, got)
		})
	}
}

func TestGetResponseTopic(t *testing.T) {
	tests := []struct {
		agentID       string
		requestID     string
		expectedTopic string
	}{
		{
			agentID:       "agent123",
			requestID:     "req789",
			expectedTopic: "pixconf/agent/agent123/response/req789",
		},
		{
			agentID:       "agent456",
			requestID:     "req101",
			expectedTopic: "pixconf/agent/agent456/response/req101",
		},
		{
			agentID:       "agent789",
			requestID:     "req202",
			expectedTopic: "pixconf/agent/agent789/response/req202",
		},
	}

	for _, tt := range tests {
		t.Run(tt.agentID+"_"+tt.requestID, func(t *testing.T) {
			got := GetResponseTopic(tt.agentID, tt.requestID)

			assert.Equal(t, tt.expectedTopic, got)
		})
	}
}
