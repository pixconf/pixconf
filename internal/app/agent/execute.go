package agent

import (
	"net/http"

	"github.com/urfave/cli/v2"
	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"
)

func (a *Agent) Execute(c *cli.Context) error {
	group, ctx := errgroup.WithContext(c.Context)

	group.Go(func() error {
		if err := a.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
