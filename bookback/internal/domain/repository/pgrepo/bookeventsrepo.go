package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type BookEventsRepo interface {
	Repository[entity.BookEvent]
}

type bookEventsRepo struct {
	repository[entity.BookEvent]
}

func NewBookEventsRepository(db postgres.Client) BookEventsRepo {
	return &bookEventsRepo{
		repository: repository[entity.BookEvent]{
			Name:   "BookEventsRepository",
			entity: entity.BookEvent{},
			db:     db,
		},
	}
}
