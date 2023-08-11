package secrets

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pixconf/pixconf/internal/xerror"
	"github.com/pixconf/pixconf/internal/xid"
	"github.com/pixconf/pixconf/pkg/secrets/protos"
)

type SecretURI struct {
	ID string `uri:"id" binding:"required"`
}

func (s *Secrets) apiSecretDetail(c *gin.Context) {
	var secURI SecretURI

	if err := c.ShouldBindUri(&secURI); err != nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	if !xid.IsValidSecretID(secURI.ID) {
		resp := xerror.ErrorSingle(http.StatusBadRequest, "wrong secret id format")
		c.JSON(resp.Code, resp)
		return
	}

	dbResp, err := s.db.GetSecretDetail(c.Request.Context(), secURI.ID)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	if dbResp == nil {
		resp := xerror.ErrorSingle(http.StatusNotFound, "secret not found")
		c.JSON(resp.Code, resp)
		return
	}

	resp := protos.SecretDetailResponse{
		ID:          dbResp.ID,
		Description: dbResp.Description,
		State:       dbResp.State,
		Tags:        dbResp.Tags,
		Alias:       dbResp.Alias,
		CreatedAt:   dbResp.CreatedAt.Time.Format(time.RFC3339),
	}

	if !dbResp.UpdatedAt.Time.IsZero() {
		resp.UpdatedAt = dbResp.UpdatedAt.Time.Format(time.RFC3339)
	}

	c.JSON(http.StatusOK, resp)
}
