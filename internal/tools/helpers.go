package tools

import (
	"context"
	"os/signal"
	"syscall"
)

func WaitInterrupted(ctx context.Context) error {
	ctxStop, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	<-ctxStop.Done()
	return ctxStop.Err()
}
