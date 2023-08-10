package postgres

import (
	"context"
	"fmt"

	"github.com/pixconf/pixconf/internal/xid"
)

func (c *Client) GetSecrets(ctx context.Context) ([]Secret, error) {
	var response []Secret

	rows, err := c.DB.Query(ctx, "select id, description, state, created_at, updated_at from secrets_secret")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row Secret
		if err := rows.Scan(
			&row.ID, &row.Description, &row.State, &row.CreatedAt, &row.UpdatedAt,
		); err != nil {
			return nil, err
		}

		row.ID = fmt.Sprintf("%s%s", xid.SecretIDPrefix, row.ID)

		response = append(response, row)
	}

	return response, nil
}
