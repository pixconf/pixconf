package secrets

import (
	"context"
	"net/http"
	"time"

	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/cmd/secrets/secrets/config"
	"github.com/pixconf/pixconf/cmd/secrets/secrets/postgres"
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
		ConnectURL: conf.DatabaseURL,
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

	group.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
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

	group.Go(func() error {
		err := xcmd.WaitInterrupted(ctx)
		server.Shutdown()
		return err
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
