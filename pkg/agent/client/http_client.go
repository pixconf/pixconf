package client

import (
	"context"
	"net"
	"net/http"
	"time"
)

func newHTTPClient(unixSocketPath string) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", unixSocketPath)
			},
		},
		Timeout: 5 * time.Second,
	}
}
