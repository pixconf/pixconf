package mqttmsg

type AgentTelemetry struct {
	AgentID     string `json:"agent_id"`
	StartedTime int64  `json:"started_time"`
}
