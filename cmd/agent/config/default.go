package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rs/xid"
)

func defaults(c *Config) error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}

	if c.AgentAPISocket == "" {
		c.AgentAPISocket = getEnvOrDefault("PIXCONF_AGENT_API_SOCKET", "/var/run/pixconf.sock")
	}

	if c.AgentID == "" {
		c.AgentID = getEnvOrDefaultFunc("PIXCONF_AGENT_ID", getHostname)
	}

	if c.Server == "" {
		c.Server = getEnvOrDefaultFunc("PIXCONF_SERVER", getServer)
	}

	return nil
}

func getHostname() string {
	if host, err := os.Hostname(); err == nil {
		return strings.ToLower(host)
	}

	getHostnameFqdn := func() string {
		row, err := exec.Command("/bin/hostname", "-f").Output()
		if err == nil {
			return strings.ToLower(strings.TrimSpace(string(row)))
		}

		return ""
	}

	if fqdn := getHostnameFqdn(); fqdn != "" {
		return fqdn
	}

	return fmt.Sprintf("dynamic-%s", xid.New().String())
}

func getServer() string {
	hostname := getHostname()

	hostnameSliced := strings.SplitN(hostname, ".", 2)
	if len(hostnameSliced) == 2 {
		return fmt.Sprintf("https://pixconf.%s", hostnameSliced[1])
	}

	return "https://pixconf.local/"
}
