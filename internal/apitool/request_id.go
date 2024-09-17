package apitool

import "github.com/gin-gonic/gin"

var requestIDCtxKeys = []string{
	"slog-gin.request-id",
}

func GetRequestID(c *gin.Context) string {
	for _, key := range requestIDCtxKeys {
		if v, ok := c.Get(key); ok {
			if s, ok := v.(string); ok {
				return s
			}
		}
	}

	return c.GetHeader("X-Request-ID")
}
