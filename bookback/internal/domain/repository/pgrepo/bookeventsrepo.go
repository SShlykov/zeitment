package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

// BookEventsRepo описывает репозиторий для работы с событиями книг.
//
//go:generate mockgen -destination=../../tests/mocks/domain/repository/pgrepo/book_events_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo BookEventsRepo
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
