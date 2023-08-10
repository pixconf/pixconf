package agent

import (
	"net/http"

	"github.com/urfave/cli/v2"
	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"
)

func (a *Agent) Execute(cCtx *cli.Context) error {
	group, ctx := errgroup.WithContext(cCtx.Context)

	group.Go(func() error {
		socketPath := cCtx.String("agent-api-socket")

		if err := a.ListenAndServe(socketPath); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	group.Go(func() error {
		err := xcmd.WaitInterrupted(ctx)
		a.Shutdown()
		return err
	})

	return group.Wait()
}
