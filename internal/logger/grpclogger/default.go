package grpclogger

import "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

var DefaultOpts = []logging.Option{
	logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
}
