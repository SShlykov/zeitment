package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type BookRepo interface {
	Repository[entity.Book]
}

type bookRepo struct {
	repository[entity.Book]
}

func NewBookRepository(db postgres.Client) BookRepo {
	return &bookRepo{
		repository: repository[entity.Book]{
			Name:   "BookRepository",
			entity: entity.Book{},
			db:     db,
		},
	}
}
