package user

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

func (uss *Service) Find(ctx context.Context, in *user_v1.ListUsersRequest) (*user_v1.ListUsersResponse, error) {
	meta := uss.getPaginationMetadata(ctx, in.Options.Pagination)
	status := &user_v1.Status{Status: "ok", Message: ""}
	resp := &user_v1.ListUsersResponse{Status: status, PaginationMetadata: meta}

	if in.Options.Pagination.Page > meta.TotalPages {
		status.Message = "Запрашиваемая страница не существует"
		status.Status = StatusError

		return resp, nil
	}

	users, err := uss.repo.List(ctx, dbutils.NewPaginationWithLimitOffset(meta.Page, meta.PageSize))
	if err != nil {
		status.Message = "Пользователи не найдены"
		status.Status = StatusError

		return resp, nil
	}

	resp.Users = adapters.UserEntitiesToProto(users)
	return resp, nil
}
