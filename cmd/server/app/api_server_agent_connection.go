package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) apiServerAgentConnectionList(c *gin.Context) {

	list := app.mqtt.Clients.GetAll()

	c.JSON(http.StatusOK, list)
}
