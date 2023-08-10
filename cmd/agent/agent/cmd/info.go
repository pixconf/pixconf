package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/pixconf/pixconf/pkg/agent/client"
)

var cmdInfo = &cli.Command{
	Name:  "info",
	Usage: "Display system-wide information",
	Action: func(cCtx *cli.Context) error {
		opts := client.Options{
			APISocketPath: cCtx.String("agent-api-socket"),
		}

		client, err := client.NewClient(opts)
		if err != nil {
			return err
		}

		resp, err := client.GetInfo(cCtx.Context)
		if err != nil {
			return err
		}

		b, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			return err
		}

		fmt.Println(string(b))

		return nil
	},
}
