package agentmeta

import "github.com/vitalvas/gokit/xstrings"

type Topics struct {
	Commands string
	Health   string
}

func GetTopics(agentID string) Topics {
	templateEnv := map[string]string{
		"client_id": agentID,
	}

	return Topics{
		Commands: xstrings.SimpleTemplate("pixconf/agent/{{ client_id }}/commands", templateEnv),
		Health:   xstrings.SimpleTemplate("pixconf/agent/{{ client_id }}/health", templateEnv),
	}
}

func GetResponseTopic(agentID string, requestID string) string {
	templateEnv := map[string]string{
		"client_id":  agentID,
		"request_id": requestID,
	}

	return xstrings.SimpleTemplate("pixconf/agent/{{ client_id }}/response/{{ request_id }}", templateEnv)
}
