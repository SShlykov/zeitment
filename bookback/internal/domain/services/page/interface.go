package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, page *entity.Page) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Page, error)
	Update(ctx context.Context, id string, page *entity.Page) (*entity.Page, error)
	Delete(ctx context.Context, id string) (*entity.Page, error)
	List(ctx context.Context) ([]entity.Page, error)
	GetPagesByChapterID(ctx context.Context, chapterID string) ([]entity.Page, error)
}
