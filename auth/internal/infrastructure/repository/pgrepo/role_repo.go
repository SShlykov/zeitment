package pgrepo

import (
	"github.com/SShlykov/zeitment/auth/internal/domain/entity"
	"github.com/SShlykov/zeitment/postgres"
)

type RolesRepo interface {
	Repository[entity.Role]
}

type rolesRepo struct {
	repository[entity.Role]
}

func NewRolesRepository(db postgres.Client) RolesRepo {
	return &rolesRepo{
		repository: repository[entity.Role]{
			Name:   "RolesRepository",
			entity: entity.Role{},
			db:     db,
		},
	}
}
