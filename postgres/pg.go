package postgres

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type key string

type Logger interface {
	Warn(msg string, attrs ...any)
	Info(msg string, attrs ...any)
	Debug(msg string, attrs ...any)
	Error(msg string, attrs ...any)
}

const (
	TxKey key = "tx"
)

type Postgres struct {
	Pool   *pgxpool.Pool
	logger Logger
}

func NewDB(dbc *pgxpool.Pool, logger Logger) DB {
	return &Postgres{Pool: dbc, logger: logger}
}

func (p *Postgres) ScanSingleContext(ctx context.Context, q Query, dest interface{}, args ...interface{}) error { //nolint:gofmt
	logQuery(ctx, p.logger, q, args...)

	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *Postgres) ScanAllContext(ctx context.Context, q Query, dest interface{}, args ...interface{}) error { //nolint:gofmt
	logQuery(ctx, p.logger, q, args...)

	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *Postgres) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) { //nolint:gofmt
	logQuery(ctx, p.logger, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.Raw, args...)
	}

	return p.Pool.Exec(ctx, q.Raw, args...)
}

func (p *Postgres) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, p.logger, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.Raw, args...)
	}

	return p.Pool.Query(ctx, q.Raw, args...)
}

func (p *Postgres) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	logQuery(ctx, p.logger, q, args...)
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.Raw, args...)
	}

	res := p.Pool.QueryRow(ctx, q.Raw, args...)

	return res
}

func (p *Postgres) QueryRawContextMulti(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.Raw, args...)
	}

	return p.Pool.Query(ctx, q.Raw, args...)
}

func (p *Postgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.Pool.BeginTx(ctx, txOptions)
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.Pool.Ping(ctx)
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func logQuery(_ context.Context, logger Logger, q Query, args ...interface{}) {
	logger.Debug(
		"executing query",
		slog.String("sql", q.Name),
		slog.String("query", q.Raw),
		slog.Any("args", args),
	)
}
