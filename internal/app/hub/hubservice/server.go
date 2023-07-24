package hubservice

import (
	"github.com/pixconf/pixconf/internal/app/hub/subscribers"
	"github.com/pixconf/pixconf/internal/logger"
	"github.com/pixconf/pixconf/internal/protos"
)

type Service struct {
	protos.UnimplementedHubServiceServer

	subscribers *subscribers.Subscribers
	log         *logger.Logger
}

type Options struct {
	Log *logger.Logger
}

func New(opts Options) *Service {
	return &Service{
		subscribers: subscribers.New(),
		log:         opts.Log,
	}
}
