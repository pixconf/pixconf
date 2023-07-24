package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/pixconf/pixconf/internal/encrypt"
)

type Epoch struct {
	ID                  int64
	EncryptionType      encrypt.Type
	PrivateKeyEncrypted string
	PrivateKey          []byte
}

func (c *Client) GetEpoch(ctx context.Context, epochID int64) (*Epoch, error) {
	row := &Epoch{}

	query := "select id, private_key, encryption_type from secrets_epoch where id=$1"
	if err := c.DB.QueryRow(ctx, query, epochID).Scan(&row.ID, &row.PrivateKeyEncrypted, &row.EncryptionType); err != nil && err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	decodedKey, err := encrypt.DecodeFromString(row.PrivateKeyEncrypted)
	if err != nil {
		return nil, err
	}

	decryptedKey, err := c.Encrypt.Decrypt(decodedKey)
	if err != nil {
		return nil, err
	}

	row.PrivateKey = decryptedKey

	return row, nil
}

func (c *Client) GetCurrentEpoch(ctx context.Context, half bool) (*Epoch, error) {
	row := &Epoch{}

	ts := time.Now().Add(-c.config.RotateEpochKeyTime.Round(time.Minute)).Unix()

	if half {
		ts /= 2
	}

	query := "select id, private_key, encryption_type from secrets_epoch where id>=$1 order by id desc limit 1"
	if err := c.DB.QueryRow(ctx, query, ts).Scan(&row.ID, &row.PrivateKeyEncrypted, &row.EncryptionType); err != nil && err == pgx.ErrNoRows {
		return c.createNewEpoch(ctx)

	} else if err != nil {
		return nil, err
	}

	decodedKey, err := encrypt.DecodeFromString(row.PrivateKeyEncrypted)
	if err != nil {
		return nil, err
	}

	decryptedKey, err := c.Encrypt.Decrypt(decodedKey)
	if err != nil {
		return nil, err
	}

	row.PrivateKey = decryptedKey

	return row, nil
}

func (c *Client) createNewEpoch(ctx context.Context) (*Epoch, error) {
	key, err := encrypt.GenerateKey()
	if err != nil {
		return nil, err
	}

	encryptedKey, err := c.Encrypt.Encrypt(key)
	if err != nil {
		return nil, err
	}

	row := &Epoch{
		ID:                  time.Now().Unix(),
		PrivateKeyEncrypted: encrypt.EncodeToString(encryptedKey),
		PrivateKey:          encryptedKey,
	}

	query := "insert into secrets_epoch (id, private_key, encryption_type) values ($1, $2, $3)"
	if _, err := c.DB.Exec(ctx, query, row.ID, row.PrivateKeyEncrypted, encrypt.TypeAesGCM); err != nil {
		return nil, err
	}

	return row, nil
}
