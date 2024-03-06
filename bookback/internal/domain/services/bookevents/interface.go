package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, chapter *entity.BookEvent) (string, error)
	FindByID(ctx context.Context, id string) (*entity.BookEvent, error)
	Update(ctx context.Context, id string, chapter *entity.BookEvent) (*entity.BookEvent, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, limit uint64, offset uint64) ([]*entity.BookEvent, error)
	FindByKV(ctx context.Context, key string, value any) ([]*entity.BookEvent, error)
}
