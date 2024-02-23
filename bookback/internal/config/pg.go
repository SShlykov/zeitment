package config

import (
	"errors"
	"os"
	"time"
)

const (
	dsnEnvVar = "DSN"
)

var ErrDSNNotSet = errors.New("DSN not set")

type PGConfig interface {
	DSN() string
	MaxPingAttempts() int
	PingInterval() time.Duration
}

type pgConfig struct {
	dsn             string
	maxPingAttempts int
	pingInterval    time.Duration
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvVar)
	if dsn == "" {
		return nil, ErrDSNNotSet
	}

	return &pgConfig{dsn: dsn, maxPingAttempts: 10, pingInterval: 3 * time.Second}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
func (c *pgConfig) MaxPingAttempts() int        { return c.maxPingAttempts }
func (c *pgConfig) PingInterval() time.Duration { return c.pingInterval }
