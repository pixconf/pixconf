package main

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/logger"
	"github.com/pixconf/pixconf/internal/tools"
)

func main() {
	log := logger.New(false)
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return tools.WaitInterrupted(ctx)
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
