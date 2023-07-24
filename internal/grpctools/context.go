package grpctools

import (
	"context"
	"net"

	"google.golang.org/grpc/peer"
)

func GetClientAddr(ctx context.Context) net.Addr {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return nil
	}

	return pr.Addr
}
