package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/postgres"
)

// MapVariablesRepo описывает репозиторий для работы с переменными карты.
//
//go:generate mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/map_variables_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo MapVariablesRepo
type MapVariablesRepo interface {
	Repository[entity.MapVariable]
}

type mapVariablesRepo struct {
	repository[entity.MapVariable]
}

func NewMapVariablesRepository(db postgres.Client) MapVariablesRepo {
	return &mapVariablesRepo{
		repository: repository[entity.MapVariable]{
			Name:   "MapVariablesRepository",
			entity: entity.MapVariable{},
			db:     db,
		},
	}
}
