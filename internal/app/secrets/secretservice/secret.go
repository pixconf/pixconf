package secretservice

import (
	"context"

	"github.com/pixconf/pixconf/internal/protos"
)

func (s *Service) CreateSecret(_ context.Context, _ *protos.SecretCreateRequest) (*protos.SecretCreateResponse, error) {
	return nil, nil
}

func (s *Service) UpdateSecret(context.Context, *protos.SecretUpdateRequest) (*protos.SecretUpdateResponse, error) {
	return nil, nil
}
