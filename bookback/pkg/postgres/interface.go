package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Handler func(ctx context.Context) error

type Client interface {
	DB() DB
	Builder() squirrel.StatementBuilderType
	Close() error
}

type DB interface {
	SQLScanner
	Transactor
	PingRunner
	Close()
}

type TxManager interface {
	ReadCommitted(ctx context.Context, handler Handler) error
}

type Query struct {
	Name string
	Raw  string
}

type Transactor interface {
	BeginTx(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type SQLScanner interface {
	ContextScanner
	QueryScanner
}

type ContextScanner interface {
	ScanSingleContext(ctx context.Context, query Query, dest interface{}, args ...interface{}) error
	ScanAllContext(ctx context.Context, query Query, dest interface{}, args ...interface{}) error
}

type QueryScanner interface {
	ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query Query, args ...interface{}) pgx.Row
	QueryRawContextMulti(ctx context.Context, query Query, args ...interface{}) (pgx.Rows, error)
}

type PingRunner interface {
	Ping(ctx context.Context) error
}
