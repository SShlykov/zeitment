package services

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models/dbutils"
)

// SimpleRepo описывает базовый интерфейс для DI репозиториев.
type SimpleRepo[T any] interface {
	Create(ctx context.Context, ent T) (string, error)
	FindByID(ctx context.Context, id string) (T, error)
	Update(ctx context.Context, id string, ent T) (T, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, options dbutils.Pagination) ([]T, error)
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]T, error)
}
