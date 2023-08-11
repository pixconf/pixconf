package postgres

import (
	"context"
	"sort"

	"github.com/jackc/pgx/v5"

	"github.com/pixconf/pixconf/internal/xid"
	"github.com/pixconf/pixconf/pkg/secrets/protos"
)

func (c *Client) CreateSecret(ctx context.Context, request protos.SecretCreateRequest) (string, error) {
	secretID, err := xid.GenerateSecretID()
	if err != nil {
		return "", err
	}

	conn, err := c.DB.Acquire(ctx)
	if err != nil {
		return secretID, err
	}

	defer conn.Release()

	transaction, err := c.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return secretID, err
	}

	defer func() {
		if err != nil {
			transaction.Rollback(ctx)
		} else {
			transaction.Commit(ctx)
		}
	}()

	query := "insert into secrets_secret (id, description, state) values ($1, $2, $3)"
	if _, err = transaction.Exec(ctx, query, secretID, request.Description, request.State); err != nil {
		return secretID, err
	}

	sort.Strings(request.Tags)

	queryInsertTags := "insert into secrets_secret_tags (secret_id, name) values ($1, $2)"
	for _, row := range request.Tags {
		if _, err = transaction.Exec(ctx, queryInsertTags, secretID, row); err != nil {
			return secretID, err
		}
	}

	queryInsertAlias := "insert into secrets_secret_alias (secret_id, name, global) values ($1, $2)"
	isGlobal := request.State == protos.SecretStateNormal.String()
	for row := range request.Alias {
		if _, err = transaction.Exec(ctx, queryInsertAlias, secretID, row, isGlobal); err != nil {
			return secretID, err
		}
	}

	return secretID, nil
}
