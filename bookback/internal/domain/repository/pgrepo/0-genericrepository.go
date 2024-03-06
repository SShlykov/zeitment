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

// Repository определяет обобщенный интерфейс для операций CRUD над сущностями типа T.
// T представляет тип сущности, который должен быть структурой.
type Repository[T any] interface {
	// List извлекает срез сущностей на основе лимита и смещения для пагинации.
	List(ctx context.Context, limit uint64, offset uint64) ([]*T, error)

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
	FindByKV(ctx context.Context, key string, value any) ([]*T, error)
}

type repository[T Simulated] struct {
	Name   string
	entity T
	db     postgres.Client
}

func (r *repository[T]) List(ctx context.Context, limit uint64, offset uint64) ([]*T, error) {
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

	var rows pgx.Rows
	rows, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return r.readList(rows)
}

func (r *repository[T]) Create(ctx context.Context, item *T) (string, error) {
	query, args, err := r.db.Builder().
		Insert(r.entity.TableName()).
		Columns(r.entity.InsertFields()...).
		Values(r.entity.EntityToInsertValues(&item)...).
		Prefix("RETURNING id").
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
	updateQuery := r.db.Builder().Update(r.entity.TableName())
	insertList := r.entity.EntityToInsertValues(item)
	for i, f := range r.entity.InsertFields() {
		updateQuery = updateQuery.Set(f, insertList[i])
	}

	query, args, err := updateQuery.
		Where("? = ?", r.entity.TableName()+`.id`, id).
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
		Where("? = ?", r.entity.TableName()+`.id`, id).
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
func (r *repository[T]) FindByKV(ctx context.Context, key string, value interface{}) ([]*T, error) {
	query, args, err := r.db.Builder().
		Select(r.entity.AllFields()...).
		From(r.entity.TableName()).
		Where("? = ?", key, value).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
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
		Where(r.entity.TableName()+".id = ?", id).
		Suffix("RETURNING " + strings.Join(r.entity.AllFields(), ", ")).
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

	result, ok := res.(*T)
	if !ok {
		return nil, fmt.Errorf("ошибка приведения типа в %s.Update", r.Name)
	}

	return result, nil
}

func (r *repository[T]) readList(rows pgx.Rows) ([]*T, error) {
	var result []*T
	defer rows.Close()

	for rows.Next() {
		item, err := r.entity.ReadItem(rows)
		if err != nil {
			return nil, err
		}

		typedItem, ok := item.(*T)
		if !ok {
			return nil, fmt.Errorf("не удалось привести тип элемента в readList")
		}

		result = append(result, typedItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
