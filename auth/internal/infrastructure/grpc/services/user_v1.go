package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/auth/internal/adapters"
	"github.com/SShlykov/zeitment/auth/internal/domain/entity"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"github.com/SShlykov/zeitment/postgres/dbutils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Repository interface {
	List(ctx context.Context, options dbutils.Pagination) ([]*entity.User, error)
	Create(ctx context.Context, item *entity.User) (string, error)
	Update(ctx context.Context, id string, item *entity.User) (*entity.User, error)
	HardDelete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*entity.User, error)
}

type userServiceServer struct {
	user_v1.UnimplementedUserServiceServer
	repo Repository
}

func NewUserServiceServer(repository Repository) user_v1.UserServiceServer {
	return &userServiceServer{repo: repository}
}

func (uss *userServiceServer) Create(ctx context.Context, in *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	userId, err := uss.repo.Create(ctx, adapters.ProtoToUser(in.User))
	if err != nil {
		return nil, err
	}

	var user *entity.User
	user, err = uss.repo.FindByID(ctx, userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user_v1.CreateUserResponse{Id: userId, Status: "ok", User: adapters.UserToProto(user)}, nil
}

func (uss *userServiceServer) Update(ctx context.Context, in *user_v1.UpdateUserRequest) (*user_v1.UpdateUserResponse, error) {
	user, err := uss.repo.Update(ctx, in.Id, adapters.ProtoToUser(in.User))
	if err != nil {
		return nil, err
	}

	return &user_v1.UpdateUserResponse{Id: user.ID, Status: "ok", User: adapters.UserToProto(user)}, nil
}

func (uss *userServiceServer) Delete(ctx context.Context, in *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	err := uss.repo.HardDelete(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (uss *userServiceServer) Get(ctx context.Context, in *user_v1.GetUserRequest) (*user_v1.GetUserResponse, error) {
	user, err := uss.repo.FindByID(ctx, in.Id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user_v1.GetUserResponse{Id: user.ID, Status: "ok", User: adapters.UserToProto(user)}, nil
}

func (uss *userServiceServer) Find(ctx context.Context, in *user_v1.ListUsersRequest) (*user_v1.ListUsersResponse, error) {
	pagination := dbutils.NewPaginationWithLimitOffset(in.Options.Pagination.Page, in.Options.Pagination.PageSize)
	users, err := uss.repo.List(ctx, pagination)
	if err != nil {
		return nil, errors.New("users not found")
	}

	return &user_v1.ListUsersResponse{Status: "ok", Users: adapters.UsersToProto(users)}, nil
}
