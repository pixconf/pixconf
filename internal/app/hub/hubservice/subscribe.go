package hubservice

import (
	"fmt"

	"github.com/pixconf/pixconf/internal/app/hub/subscribers"
	"github.com/pixconf/pixconf/internal/protos"
)

func (s *Service) Subscribe(stream protos.HubService_SubscribeServer) error {
	subs := &subscribers.Subscriber{
		ID:     "CHANGE",
		Stream: stream,
	}

	s.subscribers.Add(subs)
	defer s.subscribers.Delete(subs.ID)

	for {
		select {
		case <-stream.Context().Done():
			s.log.Debugf("client %s connection closed", subs.ID)
			return nil

		case <-subs.Close:
			s.log.Debugf("Closing stream for client ID: %s", subs.ID)

		default:
			// Default case is to avoid blocking
		}

		// Process the received message from agent
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		fmt.Printf("Received message: %#v\n", msg)
	}
}
