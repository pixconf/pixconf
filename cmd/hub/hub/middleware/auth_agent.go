package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pixconf/pixconf/internal/xerror"
)

const (
	HeaderAgentName   = "x-pixconf-agent-name"
	HeaderTimestamp   = "x-pixconf-timestamp"
	HeaderSign        = "x-pixconf-sign"
	HeaderSignHeaders = "x-pixconf-sign-headers"
)

var agentRequestHeadersRequired = []string{
	HeaderAgentName,
	HeaderTimestamp,
	HeaderSign,
	HeaderSignHeaders,
}

func AuthAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, row := range agentRequestHeadersRequired {
			if data := c.GetHeader(row); len(data) < 5 {
				resp := xerror.ErrorSingle(http.StatusUnauthorized, fmt.Sprintf("no valid '%s' header", row))
				c.JSON(resp.Code, resp)
				c.Abort()
				return
			}
		}

		signHeaders := strings.Split(c.GetHeader(HeaderSignHeaders), ";")
		if signHeaders == nil {
			resp := xerror.ErrorSingle(http.StatusUnauthorized, "no valid 'x-pix-sign-headers' header")
			c.JSON(resp.Code, resp)
			c.Abort()
			return
		}

		hash := bytes.NewBuffer(nil)
		hash.WriteString(c.Request.Method)
		hash.WriteString(c.Request.URL.Path)

		for _, row := range signHeaders {
			data := c.GetHeader(row)
			if len(data) < 3 {
				resp := xerror.ErrorSingle(http.StatusUnauthorized, fmt.Sprintf("no valid '%s' sign header", row))
				c.JSON(resp.Code, resp)
				c.Abort()
				return
			}

			if _, err := hash.Write([]byte(data)); err != nil {
				resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
				c.JSON(resp.Code, resp)
				c.Abort()
				return
			}
		}

		// name := strings.ToLower(c.GetHeader("x-pixconf-agent-name"))

		// TODO: add validate sign

		c.Next()
	}
}
