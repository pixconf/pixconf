package secrets

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/pixconf/pixconf/internal/app/secrets/protos"
	"github.com/pixconf/pixconf/internal/encrypt"
	"github.com/pixconf/pixconf/internal/xerror"
)

func (s *Secrets) apiTransitDecrypt(c *gin.Context) {
	var request protos.TransitDecryptRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	if !strings.HasPrefix(request.Data, "secrets:v1:") {
		resp := xerror.ErrorSingle(http.StatusBadRequest, "wrong encrypted format")
		c.JSON(resp.Code, resp)
		return
	}

	splited := strings.Split(request.Data, ":")
	if len(splited) != 4 {
		resp := xerror.ErrorSingle(http.StatusBadRequest, "wrong encrypted format")
		c.JSON(resp.Code, resp)
		return
	}

	epoch, err := strconv.ParseInt(splited[2], 16, 64)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	rawData, err := encrypt.DecodeFromString(splited[3])
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	epochInfo, err := s.db.GetEpoch(c.Request.Context(), epoch)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	if epochInfo == nil {
		resp := xerror.ErrorSingle(http.StatusBadRequest, "epoch not found")
		c.JSON(resp.Code, resp)
		return
	}

	enc, err := encrypt.New(epochInfo.PrivateKey, epochInfo.EncryptionType)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	decrypted, err := enc.Decrypt(rawData)
	if err != nil {
		resp := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		c.JSON(resp.Code, resp)
		return
	}

	response := protos.TransitDecryptResponse{
		Data: string(decrypted),
	}

	c.JSON(http.StatusOK, response)
}
