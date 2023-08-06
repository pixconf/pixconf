package agent

import "github.com/pixconf/pixconf/internal/logger"

type Agent struct {
	log *logger.Logger
}

type Options struct {
	Log *logger.Logger
}

func New(opts Options) *Agent {
	return &Agent{
		log: opts.Log,
	}
}
