package agent

import (
	"context"

	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/logger"
)

func Execute() {
	log := logger.New(false)
	group, ctx := errgroup.WithContext(context.Background())

	app := New()

	group.Go(func() error {
		return app.ListenAndServe()
	})

	group.Go(func() error {
		return xcmd.WaitInterrupted(ctx)
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
