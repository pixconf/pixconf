package secrets

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pixconf/pixconf/internal/xerror"
	"github.com/pixconf/pixconf/pkg/secrets/protos"
)

func (s *Secrets) apiSecretList(c *gin.Context) {
	rows, err := s.db.GetSecrets(c.Request.Context())
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	resp := protos.SecretListResponse{}

	for _, row := range rows {
		item := protos.SecretListItem{
			ID:          row.ID,
			Description: row.Description,
			State:       row.State,
			CreatedAt:   row.CreatedAt.Time.Format(time.RFC3339),
		}

		if !row.UpdatedAt.Time.IsZero() {
			item.UpdatedAt = row.UpdatedAt.Time.Format(time.RFC3339)
		}

		resp.Secrets = append(resp.Secrets, item)
	}

	c.JSON(http.StatusOK, resp)
}
