package agent

import "github.com/vitalvas/gokit/xstrings"

type mqttTopics struct {
	Commands string
	Health   string
}

func getMQTTTopics(agentID string) mqttTopics {
	templateEnv := map[string]string{
		"client_id": agentID,
	}

	return mqttTopics{
		Commands: xstrings.SimpleTemplate("pixconf/agent/{{ client_id }}/commands", templateEnv),
		Health:   xstrings.SimpleTemplate("pixconf/agent/{{ client_id }}/health", templateEnv),
	}
}
