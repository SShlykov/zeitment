package pg

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type key string

const (
	TxKey key = "tx"
)

type pg struct {
	db *pgxpool.Pool
}

func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{db: dbc}
}

func (p *pg) ScanSingleContext(ctx context.Context, q db.Query, dest interface{}, args ...interface{}) error { //nolint:gofmt
	// TODO: logQuery(ctx, q, args...)

	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAllContext(ctx context.Context, q db.Query, dest interface{}, args ...interface{}) error { //nolint:gofmt
	// TODO: logQuery(ctx, q, args...)

	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	// TODO: logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.Raw, args...)
	}

	return p.db.Exec(ctx, q.Raw, args...)
}

func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	// TODO: logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.Raw, args...)
	}

	return p.db.Query(ctx, q.Raw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	// TODO: logQuery(ctx, q, args...)
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.Raw, args...)
	}

	res := p.db.QueryRow(ctx, q.Raw, args...)

	return res
}

func (p *pg) QueryRawContextMulti(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.Raw, args...)
	}

	return p.db.Query(ctx, q.Raw, args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.db.BeginTx(ctx, txOptions)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.db.Ping(ctx)
}

func (p *pg) Close() {
	p.db.Close()
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

// func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
//	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
//	log.Println(
//		ctx,
//		fmt.Sprintf("sql: %s", q.Name),
//		fmt.Sprintf("query: %s", prettyQuery),
//	)
//}
