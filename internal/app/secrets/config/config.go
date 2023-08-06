package config

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/pixconf/pixconf/internal/encrypt"
)

type Config struct {
	APIAddress          string        `split_words:"true" default:"[::]:8142"`
	MasterEncryptionKey string        `split_words:"true" required:"true"`
	DatabaseURL         string        `split_words:"true" default:"postgres://pixconf:pixconf@localhost:5432/pixconf?sslmode=prefer"`
	RotateEpochKeyTime  time.Duration `split_words:"true" default:"336h"`
	TLSCertPath         string        `split_words:"true" required:"true"`
	TLSKeyPath          string        `split_words:"true" required:"true"`
}

func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("secrets", &c); err != nil {
		return nil, err
	}

	if c.RotateEpochKeyTime < 24*time.Hour {
		return nil, errors.New("rotate epoch key time is too short - minimum 1 day")
	}

	if key, err := base64.StdEncoding.DecodeString(c.MasterEncryptionKey); err != nil {
		return nil, err
	} else if len(key) != encrypt.KeySize {
		return nil, encrypt.ErrKeySize
	}

	return &c, nil
}
