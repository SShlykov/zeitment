package pgrepo

import (
	"context"
	"github.com/SShlykov/zeitment/postgres"
	"github.com/SShlykov/zeitment/postgres/dbutils"
	"github.com/jackc/pgx/v5"
	"strings"
)

const Equals = " = ?"

type Simulated[T any] interface {
	TableName() string
	AllFields() []string
	InsertOrUpdateFields() []string
	EntityToInsertValues(impl *T) []interface{}
	ReadItem(row pgx.Row) (T, error)
	ReadList(rows pgx.Rows) ([]T, error)
}

// Repository определяет обобщенный интерфейс для операций CRUD над сущностями типа T.
// T представляет тип сущности, который должен быть структурой.
type Repository[T any] interface {
	// Count возвращает количество сущностей в репозитории.
	Count(ctx context.Context) (uint64, error)

	// List извлекает срез сущностей на основе лимита и смещения для пагинации.
	List(ctx context.Context, options dbutils.Pagination) ([]*T, error)

	// Create вставляет новую сущность в репозиторий и возвращает её ID или ошибку.
	Create(ctx context.Context, item *T) (string, error)

	// Update изменяет существующую сущность, идентифицированную по ID. Возвращает обновленную сущность или ошибку.
	Update(ctx context.Context, id string, item *T) (*T, error)

	// HardDelete удаляет сущность, идентифицированную по ID, из репозитория. Может возвращать ошибку, если удаление не
	// удалось или сущность не найдена.
	HardDelete(ctx context.Context, id string) error

	// FindByID извлекает сущность по её ID. Возвращает указатель на сущность или ошибку.
	FindByID(ctx context.Context, id string) (*T, error)

	// FindByKV ищет сущности, соответствующие определенной паре ключ-значение.
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*T, error)
}

type repository[T Simulated[T]] struct {
	Name   string
	entity T
	db     postgres.Client
}

func (r *repository[T]) Count(ctx context.Context) (uint64, error) {
	query, args, err := r.db.Builder().
		Select("COUNT(*)").
		From(r.entity.TableName()).
		ToSql()

	if err != nil {
		return 0, err
	}

	q := postgres.Query{Name: r.Name + ".Count", Raw: query}
	var count uint64
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *repository[T]) List(ctx context.Context, options dbutils.Pagination) ([]*T, error) {
	limit, offset := options.GetPagination()
	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: r.Name + ".List", Raw: query}

	var rows pgx.Rows
	rows, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return r.readList(rows)
}

func (r *repository[T]) Create(ctx context.Context, item *T) (string, error) {
	fields := r.entity.InsertOrUpdateFields()
	insertValues := r.entity.EntityToInsertValues(item)

	query, args, err := r.db.Builder().
		Insert(r.entity.TableName()).
		Columns(fields...).
		Values(insertValues...).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return "", err
	}
	q := postgres.Query{Name: r.Name + ".Insert", Raw: query}

	var id string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository[T]) Update(ctx context.Context, id string, item *T) (*T, error) {
	insertList := r.entity.EntityToInsertValues(item)

	updateQuery := r.db.Builder().
		Update(r.entity.TableName())

	for i, field := range r.entity.InsertOrUpdateFields() {
		updateQuery = updateQuery.Set(field, insertList[i])
	}

	query, args, err :=
		updateQuery.
			Where(r.tableKey("id")+Equals, id).
			Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
			ToSql()

	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: r.Name + ".Update", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, args...)

	return r.readItem(row)
}

// HardDelete is a hard delete; But it expects that all children are deleted before
func (r *repository[T]) HardDelete(ctx context.Context, id string) error {
	query, args, err := r.db.Builder().
		Delete(r.entity.TableName()).
		Where(r.tableKey("id")+Equals, id).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return err
	}

	q := postgres.Query{Name: r.Name + ".Delete", Raw: query}
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return err
	}
	return nil
}

// FindByKV ONLY FOR QUERY COMPOSITION; KEY WILL NOT BE ESCAPED/CHECKED SO IT'S UNSAFE
// e.g. you need a function FindByTitle(title string) ([]Book, error) -> FindByKV("title", title)
func (r *repository[T]) FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*T, error) {
	key, value := options.GetFilter()
	limit, offset := options.GetPagination()

	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Where(r.tableKey(key)+Equals, value).
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: r.Name + ".FindByKV", Raw: query}

	var rows pgx.Rows
	rows, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return r.readList(rows)
}

func (r *repository[T]) FindByID(ctx context.Context, id string) (*T, error) {
	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Where(r.tableKey("id")+Equals, id).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: r.Name + ".FindByID", Raw: query}
	row := r.db.DB().QueryRowContext(ctx, q, args...)
	return r.readItem(row)
}

func (r *repository[T]) readItem(row pgx.Row) (*T, error) {
	res, err := r.entity.ReadItem(row)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repository[T]) readList(rows pgx.Rows) ([]*T, error) {
	var result []*T
	defer rows.Close()

	for rows.Next() {
		item, err := r.entity.ReadItem(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository[T]) tableKey(key string) string {
	return r.entity.TableName() + "." + key
}
