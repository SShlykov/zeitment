package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, page *entity.Page) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Page, error)
	Update(ctx context.Context, id string, page *entity.Page) (*entity.Page, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, limit uint64, offset uint64) ([]*entity.Page, error)
	FindByKV(ctx context.Context, key string, value any) ([]*entity.Page, error)
}
