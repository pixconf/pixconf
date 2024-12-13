package agent

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/pixconf/pixconf/internal/buildinfo"
)

func (app *Agent) ListenAndServe(apiSocket string) error {
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

	handler := app.apiRouterEngine()

	app.apiServer = &http.Server{
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return app.apiServer.Serve(listener)
}

func (app *Agent) apiRouterEngine() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/v1/info", app.apiInfoHandler)

	return r
}

func (app *Agent) apiInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"pid":     os.Getpid(),
		"version": buildinfo.Version,
	}

	if ppid := os.Getppid(); ppid > 0 {
		response["ppid"] = ppid
	}

	json.NewEncoder(w).Encode(response)
}
