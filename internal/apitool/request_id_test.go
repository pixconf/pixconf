package apitool

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetRequestID(t *testing.T) {
	tests := []struct {
		name          string
		contextValues map[string]interface{}
		headerValue   string
		expected      string
	}{
		{
			name: "Request ID in context",
			contextValues: map[string]interface{}{
				"slog-gin.request-id": "context-id",
			},
			headerValue: "",
			expected:    "context-id",
		},
		{
			name:          "Request ID in header",
			contextValues: map[string]interface{}{},
			headerValue:   "header-id",
			expected:      "header-id",
		},
		{
			name:          "Request ID not found",
			contextValues: map[string]interface{}{},
			headerValue:   "",
			expected:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(nil)
			for key, value := range tt.contextValues {
				c.Set(key, value)
			}
			c.Request = &http.Request{
				Header: make(http.Header),
			}
			c.Request.Header.Set("X-Request-ID", tt.headerValue)

			result := GetRequestID(c)
			require.Equal(t, tt.expected, result)
		})
	}
}
