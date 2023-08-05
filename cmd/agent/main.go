package main

import (
	"context"

	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"

	"github.com/pixconf/pixconf/internal/logger"
)

func main() {
	log := logger.New(false)
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return xcmd.WaitInterrupted(ctx)
	})

	if err := group.Wait(); err != nil {
		log.Error(err)
	}
}
