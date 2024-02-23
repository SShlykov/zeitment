package pg

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	db db.DB
}

func (c *pgClient) QueryRowContext(ctx context.Context, query string, args ...interface{}) { //nolint:gofmt
	c.db.QueryRowContext(ctx, db.Query{Name: "Query", Raw: query}, args...)
}

func NewClient(ctx context.Context, dsn string) (db.Client, error) {
	op := "pg.NewClient"
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.New(op + ": " + err.Error())
	}

	return &pgClient{db: &pg{dbc}}, nil
}

func (c *pgClient) DB() db.DB {
	return c.db
}
func (c *pgClient) Close() error {
	if c.db != nil {
		c.db.Close()
	}

	return nil
}
