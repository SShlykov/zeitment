package chapter

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, chapter *entity.Chapter) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Chapter, error)
	Update(ctx context.Context, id string, chapter *entity.Chapter) (*entity.Chapter, error)
	Delete(ctx context.Context, id string) (*entity.Chapter, error)
	List(ctx context.Context) ([]entity.Chapter, error)
	GetChapterByBookID(ctx context.Context, bookID string) ([]entity.Chapter, error)
}
