package app

import (
	"net/http"
	"time"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func (app *App) initAPI() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(location.Default())
	router.Use(sloggin.New(app.logger.With("service", "api")))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// server api
	{
		serverAPI := router.Group("/api/v1")

		serverAPI.GET("/agent/connection", app.apiServerAgentConnectionList)

		serverAPI.POST("/agent/send/command", app.apiServerAgentSendCommand)
	}

	// agent api
	{
		router.GET(".well-known/pixconf/agent-configuration", app.apiAgentAutoConfiguration)
		// agentAPI := router.Group("/api/agent/v1")

	}

	return router
}

func (app *App) ListenAndServeAPI() error {
	router := app.initAPI()

	app.apiServer = &http.Server{
		Addr:           "[::]:8080",
		Handler:        router.Handler(),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.logger.With("service", "api").Info("starting API server", "address", app.apiServer.Addr)

	if err := app.apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
