package pgrepo

import (
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres"
)

type UsersRepo interface {
	Repository[entity.User]
}

type usersRepo struct {
	repository[entity.User]
}

func NewUsersRepository(db postgres.Client) UsersRepo {
	return &usersRepo{
		repository: repository[entity.User]{
			Name:   "UsersRepository",
			entity: entity.User{},
			db:     db,
		},
	}
}
