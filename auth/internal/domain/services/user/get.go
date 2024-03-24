package user

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
)

func (uss *Service) Get(ctx context.Context, in *user_v1.GetUserRequest) (*user_v1.GetUserResponse, error) {
	user, err := uss.repo.FindByID(ctx, in.Id)
	status := &user_v1.Status{Status: "ok", Message: ""}
	resp := &user_v1.GetUserResponse{Status: status}

	if err != nil {
		status.Message = "Пользователь не найден"
		status.Status = "error"

		return resp, nil
	}

	resp.Id = user.ID
	resp.User = adapters.UserEntityToProto(user)
	return resp, nil
}
