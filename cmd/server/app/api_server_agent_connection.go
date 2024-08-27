package app

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiServerAgentConnectionListResponse struct {
	Name             string `json:"name"`
	ConnectedAddress string `json:"connected_address,omitempty"`
}

func (app *App) apiServerAgentConnectionList(c *gin.Context) {
	allClients := app.mqtt.Clients

	response := make([]apiServerAgentConnectionListResponse, 0, allClients.Len())

	for name, row := range allClients.GetAll() {
		if row == nil || row.Net.Conn == nil || row.Net.Inline {
			continue
		}

		// skip internal names
		if name != row.ID {
			continue
		}

		respRow := apiServerAgentConnectionListResponse{
			Name: name,
		}

		if remoteAddr, _, err := net.SplitHostPort(row.Net.Conn.RemoteAddr().String()); err == nil {
			respRow.ConnectedAddress = remoteAddr
		}

		response = append(response, respRow)
	}

	c.JSON(http.StatusOK, response)
}
