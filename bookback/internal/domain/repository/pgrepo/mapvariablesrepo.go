package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

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
