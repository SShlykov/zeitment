package pgrepo

import (
	"context"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"strings"
)

type Simulated interface {
	TableName() string
	AllFields() []string
	InsertFields() []string
	EntityToInsertValues(entity any) []interface{}
	ReadItem(row pgx.Row) (any, error)
	ReadList(rows pgx.Rows) ([]any, error)
}

type Repository[T Simulated] interface {
	Create(item T) (string, error)
	Update(id string, item T) (T, error)
	HardDelete(id string) (T, error)

	FindByID(id string) (T, error)
	FindByKV(key string, value interface{}) ([]T, error)

	List(ctx context.Context, limit uint64, offset uint64) ([]T, error)
}

type repository[T Simulated] struct {
	Name   string
	entity T
	db     postgres.Client
	ctx    context.Context
}

func (r *repository[T]) List(ctx context.Context, limit uint64, offset uint64) ([]T, error) {
	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Offset(offset).
		Limit(limit).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: r.Name + ".List", Raw: query}
	rows, err := r.db.DB().QueryContext(ctx, q, args...)

	return r.readList(rows)
}

func (r *repository[T]) Create(entity T) (string, error) {
	query, args, err := r.db.Builder().
		Insert(entity.TableName()).
		Columns(entity.InsertFields()...).
		Values(entity.EntityToInsertValues(entity)...).
		Prefix("RETURNING id").
		ToSql()

	if err != nil {
		return "", err
	}

	q := postgres.Query{Name: r.Name + ".Insert", Raw: query}

	var id string
	if err = r.db.DB().QueryRowContext(r.ctx, q, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository[T]) Update(id string, entity T) (T, error) {
	var zero T

	updateQuery := r.db.Builder().Update(entity.TableName())
	insertList := entity.EntityToInsertValues(entity)
	for i, f := range entity.InsertFields() {
		updateQuery = updateQuery.Set(f, insertList[i])
	}

	query, args, err := updateQuery.
		Where("? = ?", r.entity.TableName()+`.id`, id).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
		ToSql()

	if err != nil {
		return zero, err
	}

	q := postgres.Query{Name: r.Name + ".Update", Raw: query}
	row := r.db.DB().QueryRowContext(r.ctx, q, args...)

	return r.readItem(row)
}

// HardDelete is a hard delete; But it expects that all children are deleted before
func (r *repository[T]) HardDelete(id string) (T, error) {
	var zero T

	query, args, err := r.db.Builder().
		Delete(r.entity.TableName()).
		Where("? = ?", r.entity.TableName()+`.id`, id).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
		ToSql()

	if err != nil {
		return zero, err
	}

	q := postgres.Query{Name: r.Name + ".Delete", Raw: query}
	row := r.db.DB().QueryRowContext(r.ctx, q, args...)
	return r.readItem(row)
}

// FindByKV ONLY FOR QUERY COMPOSITION; KEY WILL NOT BE ESCAPED/CHECKED SO IT'S UNSAFE
// e.g. you need a function FindByTitle(title string) ([]Book, error) -> FindByKV("title", title)
func (r *repository[T]) FindByKV(key string, value interface{}) ([]T, error) {
	zero := make([]T, 0)

	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Where("? = ?", key, value).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
		ToSql()

	if err != nil {
		return zero, err
	}

	q := postgres.Query{Name: r.Name + ".FindByKV", Raw: query}
	rows, err := r.db.DB().QueryContext(r.ctx, q, args...)
	if err != nil {
		return zero, err
	}

	return r.readList(rows)
}

func (r *repository[T]) FindByID(id string) (T, error) {
	var zero T

	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Where(r.entity.TableName()+".id = ?", id).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
		ToSql()

	if err != nil {
		return zero, err
	}

	q := postgres.Query{Name: r.Name + ".FindByID", Raw: query}
	row := r.db.DB().QueryRowContext(r.ctx, q, args...)
	return r.readItem(row)
}

func (r *repository[T]) readItem(row pgx.Row) (T, error) {
	var zero T
	res, err := r.entity.ReadItem(row)
	if err != nil {
		return zero, err
	}

	result, ok := res.(T)
	if !ok {
		return zero, fmt.Errorf("ошибка приведения типа в %s.Update", r.Name)
	}

	return result, nil
}

func (r *repository[T]) readList(rows pgx.Rows) ([]T, error) {
	zero := make([]T, 0)
	res, err := r.entity.ReadList(rows)
	if err != nil {
		return zero, err
	}

	result := make([]T, len(res))
	for i, r := range res {
		result[i] = r.(T)
	}

	return result, nil
}
