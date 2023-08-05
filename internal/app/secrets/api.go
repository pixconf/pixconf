package secrets

import "github.com/gin-gonic/gin"

func (s *Secrets) routerEngine() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api/v1/secrets")

	api.POST("/transit/encrypt", s.apiTransitEncrypt)
	api.POST("/transit/decrypt", s.apiTransitDecrypt)

	return r
}
