package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type ChapterRepo interface {
	Repository[entity.Chapter]
}

type chapterRepo struct {
	repository[entity.Chapter]
}

func NewChapterRepository(db postgres.Client) ChapterRepo {
	return &chapterRepo{
		repository: repository[entity.Chapter]{
			Name:   "ChapterRepository",
			entity: entity.Chapter{},
			db:     db,
		},
	}
}
