package pgrepo

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type ParagraphRepo interface {
	Repository[entity.Paragraph]
}

type paragraphRepo struct {
	repository[entity.Paragraph]
}

func NewParagraphRepository(db postgres.Client, ctx context.Context) ParagraphRepo {
	return &paragraphRepo{
		repository: repository[entity.Paragraph]{
			Name:   "ParagraphRepository",
			entity: entity.Paragraph{},
			ctx:    ctx,
			db:     db,
		},
	}
}
