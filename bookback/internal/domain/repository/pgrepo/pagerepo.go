package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

// PageRepo описывает репозиторий для работы с страницами.
//
//go:generate mockgen -destination=../../tests/mocks/domain/repository/pgrepo/page_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo PageRepo
type PageRepo interface {
	Repository[entity.Page]
}

type pageRepo struct {
	repository[entity.Page]
}

func NewPageRepository(db postgres.Client) PageRepo {
	return &pageRepo{
		repository: repository[entity.Page]{
			Name:   "PageRepository",
			entity: entity.Page{},
			db:     db,
		},
	}
}
