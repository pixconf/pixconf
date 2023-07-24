package secretservice

import (
	"github.com/pixconf/pixconf/internal/app/secrets/postgres"
	"github.com/pixconf/pixconf/internal/protos"
)

type Service struct {
	protos.UnimplementedSecretsServer

	db *postgres.Client
}

func New(db *postgres.Client) *Service {
	return &Service{
		db: db,
	}
}
