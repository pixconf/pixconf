package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/pixconf/pixconf/pkg/secrets/protos"
)

type Secret struct {
	ID          string
	Description string
	State       string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Tags        []string
	Alias       map[string]protos.SecretAlias
}
