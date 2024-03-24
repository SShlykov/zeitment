package user

import (
	"context"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uss *Service) Delete(ctx context.Context, in *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	err := uss.repo.HardDelete(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
