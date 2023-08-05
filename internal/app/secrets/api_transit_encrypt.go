package secrets

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pixconf/pixconf/internal/app/secrets/protos"
	"github.com/pixconf/pixconf/internal/encrypt"
	"github.com/pixconf/pixconf/internal/xerror"
)

func (s *Secrets) apiTransitEncrypt(c *gin.Context) {
	var request protos.TransitEncryptRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	epoch, err := s.db.GetCurrentEpoch(c.Request.Context(), true)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	encrypted, err := s.db.Encrypt.Encrypt([]byte(request.Data))
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	response := protos.TransitEncryptResponse{
		Data: fmt.Sprintf("secrets:v1:%s:%s", strconv.FormatInt(epoch.ID, 16), encrypt.EncodeToString(encrypted)),
	}

	c.JSON(http.StatusOK, response)
}
