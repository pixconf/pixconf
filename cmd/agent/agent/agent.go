package agent

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	autopahoMemory "github.com/eclipse/paho.golang/autopaho/queue/memory"
	"github.com/pixconf/pixconf/cmd/agent/config"
	"github.com/pixconf/pixconf/internal/agent/authkey"
)

type Agent struct {
	config    *config.Config
	ctx       context.Context
	log       *slog.Logger
	apiServer *http.Server

	authKey *authkey.AuthKey

	mqttConn  *autopaho.ConnectionManager
	mqttQueue *autopahoMemory.Queue

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

		startedTime: time.Now(),

		mqttQueue: autopahoMemory.New(),
	}
}

func (app *Agent) SetConfig(config *config.Config) {
	app.config = config
}

func (app *Agent) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.startedTime = time.Time{}

	if app.apiServer != nil {
		app.apiServer.Shutdown(ctx)
		app.apiServer = nil
	}

	if app.mqttConn != nil {
		app.mqttConn.Disconnect(ctx)
		app.mqttConn = nil
	}
}
