package chapter

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, chapter *entity.Chapter) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Chapter, error)
	Update(ctx context.Context, id string, chapter *entity.Chapter) (*entity.Chapter, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, limit uint64, offset uint64) ([]*entity.Chapter, error)
	FindByKV(ctx context.Context, key string, value any) ([]*entity.Chapter, error)
}
