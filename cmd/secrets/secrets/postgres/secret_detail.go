package postgres

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/pixconf/pixconf/internal/xid"
	"github.com/pixconf/pixconf/pkg/secrets/protos"
)

func (c *Client) GetSecretDetail(ctx context.Context, secretID string) (*Secret, error) {
	response := &Secret{
		ID: secretID,
	}

	secretID = strings.TrimPrefix(secretID, xid.SecretIDPrefix)

	query := "select description, state, created_at, updated_at from secrets_secret where id=$1"
	if err := c.DB.QueryRow(ctx, query, secretID).Scan(
		&response.Description, &response.State, &response.CreatedAt, &response.UpdatedAt,
	); err != nil && err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	{
		rows, err := c.DB.Query(ctx, "select name from secrets_secret_tags where secret_id=$1 order by name asc", secretID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var tag string

			if err := rows.Scan(&tag); err != nil {
				return nil, err
			}

			response.Tags = append(response.Tags, tag)
		}
	}

	{
		rows, err := c.DB.Query(ctx, "select name from secrets_secret_alias where secret_id=$1 order by name asc", secretID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var alias string

			if err := rows.Scan(&alias); err != nil {
				return nil, err
			}

			response.Alias[alias] = protos.SecretAlias{}
		}
	}

	return response, nil
}
