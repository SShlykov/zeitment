package pgrepo

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type MapVariablesRepo interface {
	Repository[entity.MapVariable]
}

type mapVariablesRepo struct {
	repository[entity.MapVariable]
}

func NewMapVariablesRepository(db postgres.Client, ctx context.Context) MapVariablesRepo {
	return &mapVariablesRepo{
		repository: repository[entity.MapVariable]{
			Name:   "MapVariablesRepository",
			entity: entity.MapVariable{},
			ctx:    ctx,
			db:     db,
		},
	}
}
