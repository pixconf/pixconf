package secrets

import (
	"errors"
)

func (s *Secrets) ListenAndServe() error {
	if s.config == nil {
		return errors.New("no config defined")
	}

	s.log.Infof("listen on %s", s.config.APIAddr)

	router := s.routerEngine()

	return router.RunTLS(s.config.APIAddr, s.config.TLSCertPath, s.config.TLSKeyPath)
}
