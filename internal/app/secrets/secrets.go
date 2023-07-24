package secrets

import (
	"context"

	"google.golang.org/grpc"

	"github.com/pixconf/pixconf/internal/app/secrets/config"
	"github.com/pixconf/pixconf/internal/app/secrets/postgres"
	"github.com/pixconf/pixconf/internal/logger"
)

type Options struct {
	Context context.Context
	Log     *logger.Logger
	Config  *config.Config
	DB      *postgres.Client
}

type Secrets struct {
	server *grpc.Server

	config *config.Config
	ctx    context.Context
	log    *logger.Logger

	db *postgres.Client
}

func New(opts Options) *Secrets {
	return &Secrets{
		config: opts.Config,
		ctx:    opts.Context,
		log:    opts.Log,
		db:     opts.DB,
	}
}

func (s *Secrets) Shutdown() {
	if s.server != nil {
		s.server.Stop()
	}

	if s.db != nil {
		s.db.Shutdown()
	}
}
