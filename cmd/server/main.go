package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/pixconf/pixconf/cmd/server/app"
	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"
)

func main() {
	loggerOptions := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, loggerOptions))

	application, err := app.New(logger)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		// Listen in the background, if an error occurs, shutdown the application
		if err := application.ListenAndServeMQTT(); err != nil {
			fmt.Println("Error starting MQTT server", err)
			application.Shutdown()
		}
		return err
	})

	group.Go(func() error {
		err := application.ListenAndServeAPI()
		application.Shutdown()
		return err
	})

	group.Go(func() error {
		err := xcmd.WaitInterrupted(ctx)
		application.Shutdown()
		return err
	})

	if err := group.Wait(); err != nil {
		logger.ErrorContext(ctx, err.Error())
	}
}
