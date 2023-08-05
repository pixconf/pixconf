package agent

import (
	"errors"
	"net"
	"net/http"
	"os"
)

const apiSocket = "/run/pixconf-agent.sock"

func (a *Agent) ListenAndServe() error {
	if _, err := os.Stat(apiSocket); !errors.Is(err, os.ErrNotExist) {
		// TODO: add check if another agent is running
		return errors.New("found api socket")
	}

	listener, err := net.Listen("unix", apiSocket)
	if err != nil {
		return err
	}

	defer listener.Close()

	if err := os.Chmod(apiSocket, 0660); err != nil {
		return err
	}

	handler := a.apiRouterEngine()

	return http.Serve(listener, handler)
}

func (a *Agent) apiRouterEngine() *http.ServeMux {
	r := http.NewServeMux()

	return r
}
