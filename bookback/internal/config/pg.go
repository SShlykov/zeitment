package config

import (
	"errors"
	"os"
)

const (
	dsnEnvVar = "DSN"
)

var ErrDSNNotSet = errors.New("DSN not set")

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvVar)
	if dsn == "" {
		return nil, ErrDSNNotSet
	}

	return &pgConfig{dsn: dsn}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
