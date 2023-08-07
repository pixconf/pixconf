package cmd

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/pixconf/pixconf/internal/app/agent"
	"github.com/pixconf/pixconf/internal/buildinfo"
	"github.com/pixconf/pixconf/internal/logger"
)

func Execute() {
	log := logger.New(false)

	agentApp := agent.New(agent.Options{
		Log: log,
	})

	cliApp := &cli.App{
		Name:    "pixconf-agent",
		Usage:   "The PixConf Agent",
		Version: buildinfo.Version,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:       "agent-api-socket",
				Value:      "/run/pixconf-agent.sock",
				Required:   true,
				HasBeenSet: true,
				EnvVars:    []string{"PIXCONF_AGENT_API_SOCKET"},
			},
			&cli.PathFlag{
				Name:       "agent-key",
				Value:      "/etc/pixconf/agent.pem",
				Required:   true,
				HasBeenSet: true,
				EnvVars:    []string{"PIXCONF_AGENT_KEY"},
			},
		},
		Commands: []*cli.Command{
			cmdInfo,
		},

		Action: agentApp.Execute,
	}

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
