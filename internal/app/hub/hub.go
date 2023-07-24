package hub

import (
	"context"

	"google.golang.org/grpc"

	"github.com/pixconf/pixconf/internal/app/hub/config"
	"github.com/pixconf/pixconf/internal/logger"
)

type Hub struct {
	server *grpc.Server
	ctx    context.Context
	log    *logger.Logger
	config *config.Config
}

func New(ctx context.Context) *Hub {
	return &Hub{
		ctx: ctx,
	}
}

func (h *Hub) SetLogger(log *logger.Logger) *Hub {
	h.log = log
	return h
}

func (h *Hub) SetConfig(conf *config.Config) *Hub {
	h.config = conf
	return h
}

func (h *Hub) Shutdown() {
	if h.server != nil {
		h.server.Stop()
	}
}
