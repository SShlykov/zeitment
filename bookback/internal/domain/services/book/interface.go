package book

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repo interface {
	Create(ctx context.Context, book *entity.Book) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Book, error)
	Update(ctx context.Context, id string, book *entity.Book) (*entity.Book, error)
	HardDelete(ctx context.Context, id string) error
	List(ctx context.Context, limit uint64, offset uint64) ([]*entity.Book, error)
}
