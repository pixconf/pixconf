package app

import (
	"net"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type apiServerAgentConnectionListResponse []apiServerAgentConnectionListResponseRow

func (m apiServerAgentConnectionListResponse) Len() int { return len(m) }
func (m apiServerAgentConnectionListResponse) Less(i, j int) bool {
	if m[i].Name == m[j].Name {
		return m[i].Name < m[j].Name
	}
	return m[i].Name < m[j].Name
}
func (m apiServerAgentConnectionListResponse) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

type apiServerAgentConnectionListResponseRow struct {
	Name             string `json:"name"`
	ConnectedAddress string `json:"connected_address,omitempty"`
	ConnectedPort    int    `json:"connected_port,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
}

func (app *App) apiServerAgentConnectionList(c *gin.Context) {
	allClients := app.mqtt.Clients

	response := make(apiServerAgentConnectionListResponse, 0, allClients.Len())

	for name, row := range allClients.GetAll() {
		if row == nil || row.Net.Conn == nil || row.Net.Inline {
			continue
		}

		// skip internal names
		if name != row.ID {
			continue
		}

		respRow := apiServerAgentConnectionListResponseRow{
			Name:     name,
			Endpoint: row.Net.Listener,
		}

		if remoteAddr, repotePort, err := net.SplitHostPort(row.Net.Conn.RemoteAddr().String()); err == nil {
			respRow.ConnectedAddress = remoteAddr
			if port, err := strconv.Atoi(repotePort); err == nil {
				respRow.ConnectedPort = port
			}
		}

		response = append(response, respRow)
	}

	sort.Sort(response)

	c.JSON(http.StatusOK, response)
}
