package services

import (
	"context"
)

type SimpleRepo[T any] interface {
	Create(ctx context.Context, ent T) (string, error)
	FindByID(ctx context.Context, id string) (T, error)
	FindByKV(ctx context.Context, key string, value interface{}) ([]T, error)
	Update(ctx context.Context, id string, ent T) (T, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, limit uint64, offset uint64) ([]T, error)
}
