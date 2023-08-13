package secrets

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/pixconf/pixconf/internal/autocert"
	"github.com/pixconf/pixconf/internal/buildinfo"
)

func (s *Secrets) routerEngine() (*gin.Engine, error) {
	if buildinfo.Version != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	pprof.Register(r)

	if err := r.SetTrustedProxies([]string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"100.64.0.0/10",
	}); err != nil {
		return nil, err
	}

	api := r.Group("/api/v1")

	api.GET("/secrets", s.apiSecretList)
	api.POST("/secrets", s.apiSecretCreate)

	api.GET("/secrets/:id", s.apiSecretDetail)
	// api.PUT("/secrets/:id", s.apiSecretUpdate)
	// api.DELETE("/secrets/:id", s.apiSecretDelete)

	api.POST("/transit/encrypt", s.apiTransitEncrypt)
	api.POST("/transit/decrypt", s.apiTransitDecrypt)

	return r, nil
}

func (s *Secrets) ListenAndServe() error {
	if s.config == nil {
		return errors.New("no config defined")
	}

	s.log.Infof("listen on %s", s.config.APIAddress)

	router, err := s.routerEngine()
	if err != nil {
		return err
	}

	s.srv = &http.Server{
		Addr:              s.config.APIAddress,
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if s.config.TLSCertPath == "" && s.config.TLSKeyPath == "" {
		cert, privateKey, err := autocert.GenerateSelfSignedECDSACert("secrets")
		if err != nil {
			return err
		}

		tlsConfig, err := autocert.GetTLSConfig(cert, privateKey)
		if err != nil {
			return err
		}

		listen, err := tls.Listen("tcp", s.config.APIAddress, tlsConfig)
		if err != nil {
			return err
		}

		defer listen.Close()

		return s.srv.Serve(listen)
	}

	return s.srv.ListenAndServeTLS(s.config.TLSCertPath, s.config.TLSKeyPath)
}
