package secrets

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/app/secrets/config"
	"github.com/pixconf/pixconf/internal/app/secrets/postgres"
	"github.com/pixconf/pixconf/internal/logger"
	"github.com/pixconf/pixconf/internal/tools"
)

func Execute() {
	log := logger.New(false)

	group, ctx := errgroup.WithContext(context.Background())

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := postgres.NewDBClient(postgres.Options{
		Config:     conf,
		ConnectURL: conf.PostgresURL,
		Context:    ctx,
		Log:        log,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := New(Options{
		Config:  conf,
		Context: ctx,
		DB:      dbClient,
		Log:     log,
	})

	defer server.Shutdown()

	group.Go(func() error {
		return tools.WaitInterrupted(ctx)
	})

	group.Go(func() error {
		return server.ListenAndServe()
	})

	group.Go(func() error {
		runner := func(ctx context.Context) error {
			if err := server.RotateEpochSecrets(ctx); err != nil {
				log.Warnf("rotate epoch secrets: %v", err)
			}

			return nil
		}
		return tools.PeriodicRun(ctx, runner, time.Hour)
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
