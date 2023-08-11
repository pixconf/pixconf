package postgres

import (
	"context"
	"errors"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"

	"github.com/pixconf/pixconf/cmd/secrets/secrets/config"
	"github.com/pixconf/pixconf/cmd/secrets/secrets/postgres/migrations"
	"github.com/pixconf/pixconf/internal/encrypt"
	"github.com/pixconf/pixconf/internal/logger"
	"github.com/pixconf/pixconf/internal/logger/pgxlogger"
)

type Client struct {
	Context context.Context
	DB      *pgxpool.Pool

	config  *config.Config
	Encrypt encrypt.Encrypter
}

type Options struct {
	Config     *config.Config
	ConnectURL string
	Context    context.Context
	Log        *logger.Logger
}

func NewClient(opts Options) (*Client, error) {
	urlData, err := url.Parse(opts.ConnectURL)
	if err != nil {
		return nil, err
	}

	if urlData.Scheme != "postgres" {
		return nil, errors.New("database provider support postgres only")
	}

	poolConfig, err := pgxpool.ParseConfig(opts.ConnectURL)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxlogger.NewLogger(opts.Log),
		LogLevel: tracelog.LogLevelWarn,
	}

	if poolConfig.ConnConfig.RuntimeParams == nil {
		poolConfig.ConnConfig.RuntimeParams = make(map[string]string)
	}

	if _, ok := poolConfig.ConnConfig.RuntimeParams["application_name"]; !ok {
		poolConfig.ConnConfig.RuntimeParams["application_name"] = "pixconf-secrets"
	}

	conn, err := pgxpool.NewWithConfig(opts.Context, poolConfig)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(opts.Context); err != nil {
		return nil, err
	}

	if err := migrations.Migrate.RunMigrate(opts.Context, conn); err != nil {
		return nil, err
	}

	enc, err := encrypt.NewEncoded(opts.Config.MasterEncryptionKey, encrypt.TypeAesGCM)
	if err != nil {
		return nil, err
	}

	return &Client{
		DB:      conn,
		Context: opts.Context,

		config:  opts.Config,
		Encrypt: enc,
	}, nil
}

func (c *Client) Shutdown() {
	if c.DB != nil {
		c.DB.Close()
	}
}
