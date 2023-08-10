package agent

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/pixconf/pixconf/internal/xerror"
	"github.com/pixconf/pixconf/pkg/agent/protos"
)

func (a *Agent) apiInfo(response http.ResponseWriter, _ *http.Request) {
	resp := protos.InfoResponse{
		PID: os.Getpid(),
	}

	if err := json.NewEncoder(response).Encode(resp); err != nil {
		errSingle := xerror.ErrorSingle(http.StatusInternalServerError, err.Error())
		errResp, errMarshal := errSingle.Marshal()
		if errMarshal != nil {
			response.WriteHeader(http.StatusInternalServerError)
			a.log.Error(errMarshal.Error())
		} else {
			response.WriteHeader(errSingle.Code)
			response.Write(errResp)
		}
	}
}
