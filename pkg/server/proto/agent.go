package proto

type AgentAutoConfigurationResponse struct {
	ClientIPAddress string   `json:"client_ip_address"`
	MQTTEndpoints   []string `json:"mqtt_endpoints"`
}

type AgentRPCRequest struct {
	RequestID string            `json:"request_id"`
	Command   string            `json:"command"`
	Args      []string          `json:"args"`
	KWArgs    map[string]string `json:"kwargs"`
}
