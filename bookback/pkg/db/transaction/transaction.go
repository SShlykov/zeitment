package transaction

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/pkg/db"
	"github.com/SShlykov/zeitment/bookback/pkg/db/pg"
	"github.com/jackc/pgx/v5"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager создает новый менеджер транзакций, который удовлетворяет интерфейсу db.TxManager
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

// transaction основная функция, которая выполняет указанный пользователем обработчик в транзакции
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Join(err, errors.New("can't begin transaction"))
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Join(errors.New("panic recovered: %v"), errors.New(r.(string)))
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(err, errors.New("tx rollback failed"))
			}

			return
		}

		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Join(err, errors.New("tx commit failed"))
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Join(err, errors.New("failed executing code inside transaction"))
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
