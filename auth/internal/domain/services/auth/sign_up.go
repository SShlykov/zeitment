package auth

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/auth_v1"
)

func (as *Service) SignUp(ctx context.Context, in *auth_v1.SignUpRequest) (*auth_v1.SignUpResponse, error) {
	status := &auth_v1.Status{Status: "ok", Message: ""}
	resp := &auth_v1.SignUpResponse{Status: status}

	_, err := as.userService.Create(ctx, adapters.AuthProtoToUser(in.User), in.Password)
	if err != nil {
		status.Status = "error"
		status.Message = err.Error()

		return resp, nil
	}
	// TODO: добавить отправку письма с подтверждением регистрации
	// TODO: добавить роль по умолчанию
	// TODO: вернуть токен

	return resp, nil
}
