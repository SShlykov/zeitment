package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	List(ctx context.Context) ([]entity.Paragraph, error)
	Create(ctx context.Context, paragraph *entity.Paragraph) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Paragraph, error)
	Update(ctx context.Context, id string, paragraph *entity.Paragraph) (*entity.Paragraph, error)
	Delete(ctx context.Context, id string) (*entity.Paragraph, error)
	GetParagraphsByPageID(ctx context.Context, pageID string) ([]entity.Paragraph, error)
	GetParagraphsByChapterID(ctx context.Context, chapterID string) ([]entity.Paragraph, error)
	GetParagraphsByBookID(ctx context.Context, bookID string) ([]entity.Paragraph, error)
}
