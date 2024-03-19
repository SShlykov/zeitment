package pgrepo

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres"
)

type UsersRepo interface {
	Repository[entity.User]
	FindByLogin(ctx context.Context, login string) (*entity.User, error)
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

func (ur *usersRepo) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	query, args, err := ur.db.Builder().
		Select("*").
		From(ur.entity.TableName()).
		Where("login = ?", login).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := postgres.Query{Name: ur.repository.Name + ".FindByLogin", Raw: query}

	row := ur.db.DB().QueryRowContext(ctx, q, args...)

	var user entity.User
	user, err = ur.entity.ReadItem(row)

	return &user, err
}
