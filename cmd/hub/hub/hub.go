package hub

import (
	"context"
	"net/http"
	"time"

	"github.com/pixconf/pixconf/cmd/hub/hub/config"
	"github.com/pixconf/pixconf/internal/logger"
)

type Hub struct {
	srv    *http.Server
	ctx    context.Context
	log    *logger.Logger
	config *config.Config
}

type Options struct {
	Config  *config.Config
	Context context.Context
	Log     *logger.Logger
}

func New(opts Options) *Hub {
	return &Hub{
		ctx:    opts.Context,
		config: opts.Config,
		log:    opts.Log,
	}
}

func (h *Hub) Shutdown() {
	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	if h.srv != nil {
		h.srv.Shutdown(ctx)
	}
}
