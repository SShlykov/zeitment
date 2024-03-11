package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

// ChapterRepo описывает репозиторий для работы с главами.
//
//go:generate mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/chapter_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo ChapterRepo
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
