package adapters

import (
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UsersToProto(users []*entity.User) []*user_v1.User {
	protoUsers := make([]*user_v1.User, len(users))
	for i, user := range users {
		protoUsers[i] = UserToProto(user)
	}
	return protoUsers
}

func ProtoToUser(protoUser *user_v1.User) *entity.User {
	user := &entity.User{
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

func UserToProto(user *entity.User) *user_v1.User {
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
