package pgrepo

import (
	"github.com/SShlykov/zeitment/auth/internal/domain/entity"
	"github.com/SShlykov/zeitment/postgres"
)

type UserProfileRepo interface {
	Repository[entity.UserProfile]
}

type userProfilesRepo struct {
	repository[entity.UserProfile]
}

func NewUserProfilesRepository(db postgres.Client) UserProfileRepo {
	return &userProfilesRepo{
		repository: repository[entity.UserProfile]{
			Name:   "UserProfileRepository",
			entity: entity.UserProfile{},
			db:     db,
		},
	}
}
