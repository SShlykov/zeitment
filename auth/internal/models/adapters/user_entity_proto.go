package adapters

import (
	"database/sql"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func UserProtoToEntities(users []*user_v1.User) []*entity.User {
	entities := make([]*entity.User, len(users))
	for i, user := range users {
		entities[i] = UserProtoToEntity(user)
	}
	return entities
}

func UserEntitiesToProto(users []*entity.User) []*user_v1.User {
	protos := make([]*user_v1.User, len(users))
	for i, user := range users {
		protos[i] = UserEntityToProto(user)
	}
	return protos
}

func UserProtoToEntity(u *user_v1.User) *entity.User {
	return &entity.User{
		ID:          u.Id,
		CreatedAt:   ProtoToDt(u.CreatedAt),
		UpdatedAt:   ProtoToDt(u.UpdatedAt),
		DeletedAt:   ProtoToSqlDt(u.DeletedAt),
		LoggedAt:    ProtoToSqlDt(u.LoggedAt),
		ConfirmedAt: ProtoToSqlDt(u.ConfirmedAt),

		Login: u.Login,
		Email: ProtoToSqlString(u.Email),

		DeletedBy:        ProtoToSqlString(u.DeletedBy),
		AccessTemplateID: ProtoToSqlInt(u.AccessTemplateId),
		UpdateAfter:      ProtoToSqlInt64(u.UpdateAfter),
	}
}

func UserEntityToProto(u *entity.User) *user_v1.User {
	return &user_v1.User{
		Id:          u.ID,
		CreatedAt:   sqlDtToEntity(u.CreatedAt),
		UpdatedAt:   sqlDtToEntity(u.UpdatedAt),
		DeletedAt:   sqlNullDtToEntity(u.DeletedAt),
		LoggedAt:    sqlNullDtToEntity(u.LoggedAt),
		ConfirmedAt: sqlNullDtToEntity(u.ConfirmedAt),

		Login: u.Login,
		Email: sqlNullStringToEntity(u.Email),

		DeletedBy:        sqlNullStringToEntity(u.DeletedBy),
		AccessTemplateId: sqlNullIntToEntity(u.AccessTemplateID),
		UpdateAfter:      sqlNullInt64ToEntity(u.UpdateAfter),
	}
}

func ProtoToSqlString(str *wrapperspb.StringValue) sql.Null[string] {
	if str != nil {
		return sql.Null[string]{Valid: true, V: str.Value}
	}
	return sql.Null[string]{Valid: false}
}

func ProtoToSqlDt(dt *timestamppb.Timestamp) sql.Null[time.Time] {
	if dt != nil {
		return sql.Null[time.Time]{Valid: true, V: dt.AsTime()}
	}
	return sql.Null[time.Time]{Valid: false}
}

func ProtoToSqlInt(i *wrapperspb.Int32Value) sql.Null[int] {
	if i != nil {
		return sql.Null[int]{Valid: true, V: int(i.Value)}
	}
	return sql.Null[int]{Valid: false}
}

func ProtoToSqlInt64(i *wrapperspb.Int64Value) sql.Null[int64] {
	if i != nil {
		return sql.Null[int64]{Valid: true, V: i.Value}
	}
	return sql.Null[int64]{Valid: false}
}

func ProtoToDt(dt *timestamppb.Timestamp) time.Time {
	if dt != nil {
		return dt.AsTime()
	}
	return time.Time{}
}

func sqlNullStringToEntity(s sql.Null[string]) *wrapperspb.StringValue {
	if !s.Valid {
		return nil
	}
	return wrapperspb.String(s.V)
}

func sqlNullDtToEntity(dt sql.Null[time.Time]) *timestamppb.Timestamp {
	if !dt.Valid {
		return nil
	}
	return timestamppb.New(dt.V)
}

func sqlNullIntToEntity(i sql.Null[int]) *wrapperspb.Int32Value {
	if !i.Valid {
		return nil
	}
	return wrapperspb.Int32(int32(i.V))
}

func sqlNullInt64ToEntity(i sql.Null[int64]) *wrapperspb.Int64Value {
	if !i.Valid {
		return nil
	}
	return wrapperspb.Int64(i.V)
}

func sqlDtToEntity(dt time.Time) *timestamppb.Timestamp {
	return timestamppb.New(dt)
}
