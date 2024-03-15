package pgrepo

import (
	"github.com/SShlykov/zeitment/auth/internal/domain/entity"
	"github.com/SShlykov/zeitment/postgres"
)

type UserTokenRepo interface {
	Repository[entity.UserToken]
}

type userTokensRepo struct {
	repository[entity.UserToken]
}

func NewUserTokensRepository(db postgres.Client) UserTokenRepo {
	return &userTokensRepo{
		repository: repository[entity.UserToken]{
			Name:   "UserTokenRepository",
			entity: entity.UserToken{},
			db:     db,
		},
	}
}
