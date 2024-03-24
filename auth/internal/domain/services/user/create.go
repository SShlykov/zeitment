package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SShlykov/zeitment/auth/internal/domain/helper"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/internal/models/dto"
	"time"
)

func (uss *Service) Create(ctx context.Context, user *dto.User, password string) (*dto.User, error) {
	if uss.isUserExist(ctx, user.Login) {
		return nil, errors.New("пользователь с таким логином уже существует")
	}

	if err := helper.ValidateLogin(user.Login); err != nil {
		return nil, errors.New("логин не соответствует требованиям")
	}

	hashed, err := helper.HashPassword(password)
	if err != nil {
		return nil, errors.New("пароль не соответствует требованиям")
	}

	entity := adapters.UserToEntity(user)
	entity.PasswordHash = hashed
	entity.UpdateAfter = defaultUpdateAfter()

	if _, err = uss.repo.Create(ctx, entity); err != nil {
		return nil, errors.New("ошибка создания пользователя")
	}

	return adapters.EntityToUser(entity), nil
}

func defaultUpdateAfter() sql.Null[int64] {
	return sql.Null[int64]{Valid: true, V: int64(30 * 24 * time.Hour)}
}
