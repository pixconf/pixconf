package main

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/app/hub"
	"github.com/pixconf/pixconf/internal/app/hub/config"
	"github.com/pixconf/pixconf/internal/logger"
	"github.com/pixconf/pixconf/internal/tools"
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
		return tools.WaitInterrupted(ctx)
	})

	group.Go(func() error {
		return server.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
