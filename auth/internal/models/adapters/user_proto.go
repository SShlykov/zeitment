package adapters

import (
	"github.com/SShlykov/zeitment/auth/internal/models/dto"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/auth_v1"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UsersToProto(users []*dto.User) []*user_v1.User {
	protoUsers := make([]*user_v1.User, len(users))
	for i, user := range users {
		protoUsers[i] = UserToProto(user)
	}
	return protoUsers
}

func ProtoToUsers(protoUsers []*user_v1.User) []*dto.User {
	users := make([]*dto.User, len(protoUsers))
	for i, protoUser := range protoUsers {
		users[i] = ProtoToUser(protoUser)
	}
	return users
}

func ProtoToUser(protoUser *user_v1.User) *dto.User {
	user := &dto.User{
		ID:          protoUser.Id,
		CreatedAt:   protoUser.CreatedAt.AsTime(),
		UpdatedAt:   protoUser.UpdatedAt.AsTime(),
		DeletedAt:   ProtoToNullDt(protoUser.DeletedAt),
		LoggedAt:    ProtoToNullDt(protoUser.LoggedAt),
		ConfirmedAt: ProtoToNullDt(protoUser.ConfirmedAt),

		Login: protoUser.Login,
		Email: ProtoToNullString(protoUser.Email),

		DeletedBy:        ProtoToNullString(protoUser.DeletedBy),
		AccessTemplateID: ProtoToNullInt(protoUser.AccessTemplateId),
		UpdateAfter:      ProtoToNullInt64(protoUser.UpdateAfter),
	}

	return user
}

func UserToProto(user *dto.User) *user_v1.User {
	protoUser := &user_v1.User{
		Id:          user.ID,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
		DeletedAt:   NullDtToProto(user.DeletedAt),
		LoggedAt:    NullDtToProto(user.LoggedAt),
		ConfirmedAt: NullDtToProto(user.ConfirmedAt),
		Login:       user.Login,
		Email:       NullStringToProto(user.Email),

		DeletedBy:        NullStringToProto(user.DeletedBy),
		AccessTemplateId: NullIntToProto(user.AccessTemplateID),
		UpdateAfter:      NullInt64ToProto(user.UpdateAfter),
	}

	return protoUser
}

func AuthUserToProto(user *dto.User) *auth_v1.User {
	protoUser := &auth_v1.User{
		Id:          user.ID,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
		DeletedAt:   NullDtToProto(user.DeletedAt),
		LoggedAt:    NullDtToProto(user.LoggedAt),
		ConfirmedAt: NullDtToProto(user.ConfirmedAt),
		Login:       user.Login,
		Email:       NullStringToProto(user.Email),

		DeletedBy:        NullStringToProto(user.DeletedBy),
		AccessTemplateId: NullIntToProto(user.AccessTemplateID),
		UpdateAfter:      NullInt64ToProto(user.UpdateAfter),
	}

	return protoUser
}

func AuthProtoToUser(protoUser *auth_v1.User) *dto.User {
	user := &dto.User{
		ID:          protoUser.Id,
		CreatedAt:   protoUser.CreatedAt.AsTime(),
		UpdatedAt:   protoUser.UpdatedAt.AsTime(),
		DeletedAt:   ProtoToNullDt(protoUser.DeletedAt),
		LoggedAt:    ProtoToNullDt(protoUser.LoggedAt),
		ConfirmedAt: ProtoToNullDt(protoUser.ConfirmedAt),

		Login: protoUser.Login,
		Email: ProtoToNullString(protoUser.Email),

		DeletedBy:        ProtoToNullString(protoUser.DeletedBy),
		AccessTemplateID: ProtoToNullInt(protoUser.AccessTemplateId),
		UpdateAfter:      ProtoToNullInt64(protoUser.UpdateAfter),
	}

	return user
}
