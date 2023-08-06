package secrets

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/pixconf/pixconf/internal/buildinfo"
)

func (s *Secrets) routerEngine() (*gin.Engine, error) {
	if buildinfo.Version != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	if err := r.SetTrustedProxies([]string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"100.64.0.0/10",
	}); err != nil {
		return nil, err
	}

	api := r.Group("/api/v1/secrets")

	api.POST("/transit/encrypt", s.apiTransitEncrypt)
	api.POST("/transit/decrypt", s.apiTransitDecrypt)

	return r, nil
}

func (s *Secrets) ListenAndServe() error {
	if s.config == nil {
		return errors.New("no config defined")
	}

	s.log.Infof("listen on %s", s.config.APIAddr)

	router, err := s.routerEngine()
	if err != nil {
		return err
	}

	return router.RunTLS(s.config.APIAddr, s.config.TLSCertPath, s.config.TLSKeyPath)
}
