package secrets

import (
	"context"
	"net/http"
	"time"

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
	config *config.Config
	ctx    context.Context
	log    *logger.Logger

	db  *postgres.Client
	srv *http.Server
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
	ctx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()

	if s.srv != nil {
		s.srv.Shutdown(ctx)
	}

	if s.db != nil {
		s.db.Shutdown()
	}
}
