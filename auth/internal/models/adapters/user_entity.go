package adapters

import (
	"database/sql"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/auth/internal/models/dto"
	"github.com/SShlykov/zeitment/auth/internal/models/types"
	"time"
)

func EntitiesToUsers(users []*entity.User) []*dto.User {
	dtos := make([]*dto.User, len(users))
	for i, user := range users {
		dtos[i] = EntityToUser(user)
	}
	return dtos
}

func UserToEntity(u *dto.User) *entity.User {
	return &entity.User{
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   typesNullDtToSQL(u.DeletedAt),
		LoggedAt:    typesNullDtToSQL(u.LoggedAt),
		ConfirmedAt: typesNullDtToSQL(u.ConfirmedAt),

		Login: u.Login,
		Email: typesNullStringToSQL(u.Email),

		DeletedBy:        typesNullStringToSQL(u.DeletedBy),
		AccessTemplateID: typesNullIntToSQL(u.AccessTemplateID),
		UpdateAfter:      typesNullInt64ToSQL(u.UpdateAfter),
	}
}

func EntityToUser(u *entity.User) *dto.User {
	return &dto.User{
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   sqlNullDtToTypes(u.DeletedAt),
		LoggedAt:    sqlNullDtToTypes(u.LoggedAt),
		ConfirmedAt: sqlNullDtToTypes(u.ConfirmedAt),

		Login: u.Login,
		Email: sqlNullStringToTypes(u.Email),

		DeletedBy:        sqlNullStringToTypes(u.DeletedBy),
		AccessTemplateID: sqlNullIntToTypes(u.AccessTemplateID),
		UpdateAfter:      sqlNullInt64ToTypes(u.UpdateAfter),
	}
}

func sqlNullStringToTypes(s sql.Null[string]) types.Null[string] {
	if !s.Valid {
		return types.Null[string]{Valid: false}
	}
	return types.Null[string]{Valid: true, Value: s.V}
}

func sqlNullDtToTypes(dt sql.Null[time.Time]) types.Null[time.Time] {
	if !dt.Valid {
		return types.Null[time.Time]{Valid: false}
	}
	return types.Null[time.Time]{Valid: true, Value: dt.V}
}

func sqlNullIntToTypes(i sql.Null[int]) types.Null[int] {
	if !i.Valid {
		return types.Null[int]{Valid: false}
	}
	return types.Null[int]{Valid: true, Value: i.V}
}

func sqlNullInt64ToTypes(i sql.Null[int64]) types.Null[int64] {
	if !i.Valid {
		return types.Null[int64]{Valid: false}
	}
	return types.Null[int64]{Valid: true, Value: i.V}
}

func typesNullStringToSQL(s types.Null[string]) sql.Null[string] {
	if !s.Valid {
		return sql.Null[string]{Valid: false}
	}
	return sql.Null[string]{Valid: true, V: s.Value}
}

func typesNullDtToSQL(dt types.Null[time.Time]) sql.Null[time.Time] {
	if !dt.Valid {
		return sql.Null[time.Time]{Valid: false}
	}
	return sql.Null[time.Time]{Valid: true, V: dt.Value}
}

func typesNullIntToSQL(i types.Null[int]) sql.Null[int] {
	if !i.Valid {
		return sql.Null[int]{Valid: false}
	}
	return sql.Null[int]{Valid: true, V: i.Value}
}

func typesNullInt64ToSQL(i types.Null[int64]) sql.Null[int64] {
	if !i.Valid {
		return sql.Null[int64]{Valid: false}
	}
	return sql.Null[int64]{Valid: true, V: i.Value}
}
