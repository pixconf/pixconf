package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pixconf/pixconf/internal/xerror"
	"github.com/pixconf/pixconf/pkg/secrets/protos"
)

func (s *Secrets) apiSecretCreate(c *gin.Context) {
	var request protos.SecretCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	if request.State == "" {
		request.State = protos.SecretStateNormal.String()
	}

	if errorRows := request.Validate(); errorRows != nil {
		errors := xerror.Response{
			Code:   http.StatusBadRequest,
			Errors: errorRows,
		}
		c.JSON(errors.Code, errors)
		return
	}

	secretID, err := s.db.CreateSecret(c.Request.Context(), request)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, protos.SecretCreateResponse{
		ID: secretID,
	})
}
