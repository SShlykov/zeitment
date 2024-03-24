package user

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
)

func (uss *Service) Update(ctx context.Context, in *user_v1.UpdateUserRequest) (*user_v1.UpdateUserResponse, error) {
	status := &user_v1.Status{Status: "ok", Message: ""}
	resp := &user_v1.UpdateUserResponse{Status: status}
	user, err := uss.repo.Update(ctx, in.Id, adapters.UserProtoToEntity(in.User))

	if err != nil {
		status.Message = "Пользователь не найден"
		status.Status = "error"

		return resp, nil
	}

	resp.User = adapters.UserEntityToProto(user)
	resp.Id = user.ID
	return resp, nil
}
