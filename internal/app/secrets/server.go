package secrets

import (
	"errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

	"github.com/pixconf/pixconf/internal/app/secrets/secretservice"
	"github.com/pixconf/pixconf/internal/grpctools"
	"github.com/pixconf/pixconf/internal/logger/grpclogger"
	"github.com/pixconf/pixconf/internal/protos"
)

func (s *Secrets) ListenAndServe() error {
	if s.config == nil {
		return errors.New("no config defined")
	}

	tlsCredentials, err := grpctools.LoadTLSCredentials(s.config.TLSCertPath, s.config.TLSKeyPath)
	if err != nil {
		return err
	}

	serverOpts := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
		grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(grpclogger.InterceptorLogger(s.log), grpclogger.DefaultOpts...)),
		grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(grpclogger.InterceptorLogger(s.log), grpclogger.DefaultOpts...)),
	}

	s.server = grpc.NewServer(serverOpts...)

	reflection.Register(s.server)

	secretService := secretservice.New(s.db)

	protos.RegisterSecretsServer(s.server, secretService)

	listen, err := net.Listen("tcp", s.config.APIAddr)
	if err != nil {
		return err
	}

	s.log.Infof("listen on %s", listen.Addr().String())

	return s.server.Serve(listen)
}
