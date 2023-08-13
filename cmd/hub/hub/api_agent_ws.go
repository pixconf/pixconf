package hub

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{"pixconf-hub"},
}

func (h *Hub) apiAgentWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			h.log.Errorf("error reading message: %v", err)
			break
		}

		fmt.Printf("Received message: %s\n", p)
		if err := conn.WriteMessage(messageType, p); err != nil {
			h.log.Errorf("error writing message: %v", err)
			break
		}
	}
}
