package postgres

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/SShlykov/zeitment/bookback/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type pgClient struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	db      DB
	builder squirrel.StatementBuilderType
}

func NewClient(ctx context.Context, logger logger.Logger, dsn string) (Client, error) {
	client := &pgClient{
		builder:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		slog.Error("Cant parse dsn", slog.String("dsn", dsn))
		return nil, err
	}

	poolConfig.MaxConns = int32(_defaultMaxPoolSize)

	return client.Connect(ctx, logger, poolConfig)
}

func (c *pgClient) Connect(ctx context.Context, logger logger.Logger, poolConfig *pgxpool.Config) (Client, error) {
	logger.Debug(
		"connecting to db",
		slog.Int("attempts", c.connAttempts),
		slog.String("dsn", poolConfig.ConnString()),
		slog.Int("maxPoolSize", _defaultMaxPoolSize),
	)
	for c.connAttempts > 0 {
		var pool *pgxpool.Pool
		pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
		if err == nil {
			c.db = NewDB(pool, logger)
			return c, nil
		}

		time.Sleep(c.connTimeout)

		c.connAttempts--
	}

	return nil, errors.New("failed to connect to db")
}

func (c *pgClient) DB() DB {
	return c.db
}

func (c *pgClient) Builder() squirrel.StatementBuilderType {
	return c.builder
}

func (c *pgClient) Close() error {
	if c.db != nil {
		c.db.Close()
	}

	return nil
}
