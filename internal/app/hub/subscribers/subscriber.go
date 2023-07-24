package subscribers

import (
	"time"

	"github.com/pixconf/pixconf/internal/protos"
)

type Subscriber struct {
	ID     string
	Stream protos.HubService_SubscribeServer

	Close chan bool

	ConnectedTime time.Time // auto
}
