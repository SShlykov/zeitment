package pgrepo

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

// BookRepo описывает репозиторий для работы с книгами.
//
//go:generate mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/book_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo BookRepo
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
