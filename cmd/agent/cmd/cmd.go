package cmd

import (
	"log/slog"
	"os"

	"github.com/pixconf/pixconf/cmd/agent/agent"
	"github.com/pixconf/pixconf/internal/buildinfo"
	"github.com/urfave/cli/v2"
)

func Execute() {
	loggerOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, loggerOptions))

	agentApp := agent.New(agent.Options{
		Log: logger,
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
			&cli.StringFlag{
				Name:       "server",
				Value:      "http://localhost:8080",
				Required:   true,
				HasBeenSet: true,
				EnvVars:    []string{"PIXCONF_SERVER"},
			},
		},
		Action: agentApp.Execute,
	}

	if err := cliApp.Run(os.Args); err != nil {
		logger.Error(err.Error())
	}
}
