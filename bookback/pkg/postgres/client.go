package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres/pg"
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

func NewClient(ctx context.Context, logger *slog.Logger, dsn string) (Client, error) {
	client := &pgClient{builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		slog.Error("Cant parse dsn", slog.String("dsn", dsn))
		return nil, err
	}

	poolConfig.MaxConns = int32(_defaultMaxPoolSize)

	for client.connAttempts > 0 {
		var pool *pgxpool.Pool
		pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err == nil {
			client.db = pg.NewDB(pool, logger)
			return client, nil
		}

		time.Sleep(client.connTimeout)

		client.connAttempts--
	}

	return nil, err
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
