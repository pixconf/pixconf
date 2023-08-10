package agent

import (
	"context"
	"net/http"
	"time"

	"github.com/pixconf/pixconf/internal/logger"
)

type Agent struct {
	ctx       context.Context
	apiServer *http.Server
	log       *logger.Logger
}

type Options struct {
	Context context.Context
	Log     *logger.Logger
}

func New(opts Options) *Agent {
	return &Agent{
		log: opts.Log,
		ctx: opts.Context,
	}
}

func (a *Agent) Shutdown() {
	ctx, cancel := context.WithTimeout(a.ctx, 5*time.Second)
	defer cancel()

	if a.apiServer != nil {
		a.apiServer.Shutdown(ctx)
	}
}
