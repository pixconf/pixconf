package agent

import (
	"context"

	"github.com/urfave/cli/v2"
	"github.com/vitalvas/gokit/xcmd"
	"golang.org/x/sync/errgroup"
)

func (a *Agent) Execute(_ *cli.Context) error {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return a.ListenAndServe()
	})

	group.Go(func() error {
		return xcmd.WaitInterrupted(ctx)
	})

	return group.Wait()
}
