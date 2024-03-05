package bookevents

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, event *entity.BookEvent) (string, error)
	FindByID(ctx context.Context, id string) (*entity.BookEvent, error)
	Update(ctx context.Context, id string, event *entity.BookEvent) (*entity.BookEvent, error)
	Delete(ctx context.Context, id string) (*entity.BookEvent, error)
	GetByBookID(ctx context.Context, bookID string) ([]entity.BookEvent, error)
	GetByChapterID(ctx context.Context, chapterID string) ([]entity.BookEvent, error)
	GetByPageID(ctx context.Context, pageID string) ([]entity.BookEvent, error)
	GetByParagraphID(ctx context.Context, paragraphID string) ([]entity.BookEvent, error)
}
