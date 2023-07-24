package hub

import (
	"errors"
	"net"

	"google.golang.org/grpc"

	"github.com/pixconf/pixconf/internal/app/hub/hubservice"
	"github.com/pixconf/pixconf/internal/protos"
)

func (h *Hub) ListenAndServe() error {
	if h.config == nil {
		return errors.New("no config defined")
	}

	serverOpts := []grpc.ServerOption{}

	h.server = grpc.NewServer(serverOpts...)

	hubService := hubservice.New(hubservice.Options{
		Log: h.log,
	})

	protos.RegisterHubServiceServer(h.server, hubService)

	listen, err := net.Listen("tcp", "[::]:8140")
	if err != nil {
		return err
	}

	return h.server.Serve(listen)
}
