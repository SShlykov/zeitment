package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres"
)

// ParagraphRepo описывает репозиторий для работы с параграфами.
//
//go:generate mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/paragraph_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo ParagraphRepo
type ParagraphRepo interface {
	Repository[entity.Paragraph]
}

type paragraphRepo struct {
	repository[entity.Paragraph]
}

func NewParagraphRepository(db postgres.Client) ParagraphRepo {
	return &paragraphRepo{
		repository: repository[entity.Paragraph]{
			Name:   "ParagraphRepository",
			entity: entity.Paragraph{},
			db:     db,
		},
	}
}
