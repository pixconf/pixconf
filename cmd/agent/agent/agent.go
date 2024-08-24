package agent

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/rs/xid"
)

type Agent struct {
	ctx            context.Context
	log            *slog.Logger
	apiServer      *http.Server
	serverEndpoint string

	mqttConn     *autopaho.ConnectionManager
	mqttClientID string

	startedTime time.Time
}

type Options struct {
	Context context.Context
	Log     *slog.Logger
}

func New(opts Options) *Agent {
	return &Agent{
		ctx: opts.Context,
		log: opts.Log,

		mqttClientID: fmt.Sprintf("agent-%s", xid.New().String()),
		startedTime:  time.Now(),
	}
}

func (app *Agent) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if app.apiServer != nil {
		app.apiServer.Shutdown(ctx)
		app.apiServer = nil
	}

	if app.mqttConn != nil {
		app.mqttConn.Disconnect(ctx)
		app.mqttConn = nil
	}
}
