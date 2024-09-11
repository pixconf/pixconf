package cmd

import (
	"log/slog"
	"os"

	"github.com/pixconf/pixconf/cmd/agent/agent"
	"github.com/pixconf/pixconf/cmd/agent/config"
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
				Name:       "config",
				Value:      "/etc/pixconf/agent.yaml",
				Required:   true,
				HasBeenSet: true,
				EnvVars:    []string{"PIXCONF_AGENT_CONFIG"},
			},
		},
		Before: func(c *cli.Context) error {
			configPath := c.String("config")

			conf, err := config.Load(configPath)
			if err != nil {
				return err
			}

			agentApp.SetConfig(conf)

			return nil
		},
		Action: agentApp.Execute,
	}

	if err := cliApp.Run(os.Args); err != nil {
		logger.Error(err.Error())
	}
}
