package pgrepo

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type ChapterRepo interface {
	Repository[entity.Chapter]
}

type chapterRepo struct {
	repository[entity.Chapter]
}

func NewChapterRepository(db postgres.Client, ctx context.Context) ChapterRepo {
	return &chapterRepo{
		repository: repository[entity.Chapter]{
			Name:   "ChapterRepository",
			entity: entity.Chapter{},
			ctx:    ctx,
			db:     db,
		},
	}
}
