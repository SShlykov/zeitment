package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, chapter *entity.MapVariable) (string, error)
	FindByID(ctx context.Context, id string) (*entity.MapVariable, error)
	Update(ctx context.Context, id string, chapter *entity.MapVariable) (*entity.MapVariable, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, limit uint64, offset uint64) ([]*entity.MapVariable, error)
	FindByKV(ctx context.Context, key string, value any) ([]*entity.MapVariable, error)
}
