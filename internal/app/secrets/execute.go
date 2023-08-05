package secrets

import (
	"context"
	"time"

	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/app/secrets/config"
	"github.com/pixconf/pixconf/internal/app/secrets/postgres"
	"github.com/pixconf/pixconf/internal/logger"
)

func Execute() {
	log := logger.New(false)

	group, ctx := errgroup.WithContext(context.Background())

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := postgres.NewClient(postgres.Options{
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
		return xcmd.WaitInterrupted(ctx)
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

		return xcmd.PeriodicRun(ctx, runner, time.Hour)
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
