package agent

import (
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"
)

func (app *Agent) Execute(cliCtx *cli.Context) error {
	app.serverEndpoint = cliCtx.String("server")

	group, ctx := errgroup.WithContext(cliCtx.Context)

	group.Go(func() error {
		socketPath := cliCtx.String("agent-api-socket")

		if err := app.ListenAndServe(socketPath); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	group.Go(func() error {
		defer app.Shutdown()

		for {
			select {
			case <-ctx.Done():
				return nil

			default:
				if err := app.mqttConnect(ctx); err != nil {
					app.log.Error(err.Error())
				}

				time.Sleep(5 * time.Second)
			}

		}

		// return nil
	})

	group.Go(func() error {
		return xcmd.PeriodicRun(ctx, app.mqttSendHealthTelemetry, 3*time.Minute)
	})

	group.Go(func() error {
		err := xcmd.WaitInterrupted(ctx)
		app.Shutdown()
		return err
	})

	return group.Wait()
}
