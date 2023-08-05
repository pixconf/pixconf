package main

import (
	"context"

	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/app/hub"
	"github.com/pixconf/pixconf/internal/app/hub/config"
	"github.com/pixconf/pixconf/internal/logger"
)

func main() {
	log := logger.New(false)

	group, ctx := errgroup.WithContext(context.Background())

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	server := hub.New(ctx).SetLogger(log).SetConfig(conf)

	defer server.Shutdown()

	group.Go(func() error {
		return xcmd.WaitInterrupted(ctx)
	})

	group.Go(func() error {
		return server.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
