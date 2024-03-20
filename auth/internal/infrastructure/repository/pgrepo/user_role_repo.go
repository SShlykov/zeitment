package pgrepo

import (
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres"
)

type UserRoleRepo interface {
	Repository[entity.UserRole]
}

type userRolesRepo struct {
	repository[entity.UserRole]
}

func NewUserRolesRepository(db postgres.Client) UserRoleRepo {
	return &userRolesRepo{
		repository: repository[entity.UserRole]{
			Name:   "UserRoleRepository",
			entity: entity.UserRole{},
			db:     db,
		},
	}
}
