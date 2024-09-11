package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rs/xid"
)

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
