package hub

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/pixconf/pixconf/cmd/hub/hub/middleware"
	"github.com/pixconf/pixconf/internal/autocert"
	"github.com/pixconf/pixconf/internal/buildinfo"
)

func (h *Hub) routerEngine() (*gin.Engine, error) {
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

	api := r.Group("/api/v1/hub")

	agentAPI := api.Group("/agent")

	agentAuthed := agentAPI.Group("")
	agentAuthed.Use(middleware.AuthAgent())
	agentAuthed.GET("/ws/:agent", h.apiAgentWS)

	return r, nil
}

func (h *Hub) ListenAndServe() error {
	if h.config == nil {
		return errors.New("no config defined")
	}

	h.log.Infof("listen on %s", h.config.APIAddress)

	router, err := h.routerEngine()
	if err != nil {
		return err
	}

	h.srv = &http.Server{
		Addr:              h.config.APIAddress,
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if h.config.TLSCertPath == "" && h.config.TLSKeyPath == "" {
		cert, privateKey, err := autocert.GenerateSelfSignedECDSACert("hub")
		if err != nil {
			return err
		}

		tlsConfig, err := autocert.GetTLSConfig(cert, privateKey)
		if err != nil {
			return err
		}

		listen, err := tls.Listen("tcp", h.config.APIAddress, tlsConfig)
		if err != nil {
			return err
		}

		defer listen.Close()

		return h.srv.Serve(listen)
	}

	return h.srv.ListenAndServeTLS(h.config.TLSCertPath, h.config.TLSKeyPath)
}
