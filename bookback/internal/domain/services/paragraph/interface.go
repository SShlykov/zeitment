package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	List(ctx context.Context, limit uint64, offset uint64) ([]*entity.Paragraph, error)
	Create(ctx context.Context, paragraph *entity.Paragraph) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Paragraph, error)
	Update(ctx context.Context, id string, paragraph *entity.Paragraph) (*entity.Paragraph, error)
	HardDelete(ctx context.Context, id string) error
	FindByKV(ctx context.Context, key string, value interface{}) ([]*entity.Paragraph, error)
}
