package proto

type AgentAutoConfigurationResponse struct {
	ClientIPAddress string   `json:"client_ip_address"`
	MQTTEndpoints   []string `json:"mqtt_endpoints"`
}
