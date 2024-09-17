package proto

type AgentAutoConfigurationResponse struct {
	ClientIPAddress string   `json:"client_ip_address"`
	MQTTEndpoints   []string `json:"mqtt_endpoints"`
}

type AgentRPCRequest struct {
	RequestID string            `json:"request_id"`
	Command   string            `json:"command"`
	Args      []string          `json:"args,omitempty"`
	KWArgs    map[string]string `json:"kwargs,omitempty"`
}

type AgentRPCResponse struct {
	RequestID     string           `json:"request_id"`
	Request       *AgentRPCRequest `json:"request,omitempty"`
	ExecutionTime float64          `json:"execution_time,omitempty"` // in seconds
}
