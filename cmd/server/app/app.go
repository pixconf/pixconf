package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/pixconf/pixconf/cmd/server/config"
)

type App struct {
	apiServer *http.Server
	config    *config.Config
	logger    *slog.Logger
	mqtt      *mqtt.Server
}

func New(logger *slog.Logger) (*App, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	app := &App{
		config: conf,
		logger: logger,
	}

	if err := app.initMQTT(); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if app.mqtt != nil {
		app.mqtt.Close()
		app.mqtt = nil
	}

	if app.apiServer != nil {
		if err := app.apiServer.Shutdown(ctx); err != nil {
			app.logger.Error(err.Error())
		}

		app.apiServer = nil
	}
}
