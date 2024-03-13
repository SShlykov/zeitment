package config

import (
	"errors"
	"fmt"
	"os"
	"time"
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
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DB_NAME")
	sslMode := os.Getenv("PG_SSL_MODE")

	// Формирование DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode)
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
