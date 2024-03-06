package pgrepo

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type PageRepo interface {
	Repository[entity.Page]
}

type pageRepo struct {
	repository[entity.Page]
}

func NewPageRepository(db postgres.Client, ctx context.Context) PageRepo {
	return &pageRepo{
		repository: repository[entity.Page]{
			Name:   "PageRepository",
			entity: entity.Page{},
			ctx:    ctx,
			db:     db,
		},
	}
}
