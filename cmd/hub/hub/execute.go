package hub

import (
	"context"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/cmd/hub/hub/config"
	"github.com/pixconf/pixconf/internal/logger"
	"github.com/vitalvas/gokit/xcmd"
)

func Execute() {
	log := logger.New(false)

	group, ctx := errgroup.WithContext(context.Background())

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	server := New(Options{
		Config:  conf,
		Context: ctx,
		Log:     log,
	})

	group.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
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
