package secretservice

import (
	"context"

	"github.com/pixconf/pixconf/internal/protos"
)

func (s *Service) CreateSecretVersion(context.Context, *protos.SecretCreateVersionRequest) (*protos.SecretCreateVersionResponse, error) {
	return nil, nil
}
