package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/auth/internal/domain/adapters"
	"github.com/SShlykov/zeitment/auth/internal/domain/helper"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"github.com/SShlykov/zeitment/postgres/dbutils"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type Repository interface {
	List(ctx context.Context, options dbutils.Pagination) ([]*entity.User, error)
	Create(ctx context.Context, item *entity.User) (string, error)
	Update(ctx context.Context, id string, item *entity.User) (*entity.User, error)
	HardDelete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*entity.User, error)
	FindByLogin(ctx context.Context, login string) (*entity.User, error)
}

type userServiceServer struct {
	user_v1.UnimplementedUserServiceServer
	repo Repository
}

func NewUserServiceServer(repository Repository) user_v1.UserServiceServer {
	return &userServiceServer{repo: repository}
}

func (uss *userServiceServer) SignUp(ctx context.Context, in *user_v1.SignUpRequest) (*user_v1.SignUpResponse, error) {

	if in.User == nil || in.Password == "" {
		return nil, errors.New("пользователь или пароль не могут быть пустыми; ошибка протокола")
	}
	if uss.isUserExist(ctx, in.User.Login) {
		return nil, errors.New("пользователь с таким логином уже существует")
	}
	if err := helper.ValidateLogin(in.User.Login); err != nil {
		fmt.Println("Ошибка валидации пароля:", err)
	}
	hashed, err := helper.HashPassword(in.Password)
	if err != nil {
		fmt.Println("Ошибка обработки пароля:", err)
	}

	user := adapters.ProtoToUser(in.User)
	user.PasswordHash = hashed
	user.UpdateAfter = sql.Null[int64]{Valid: true, V: int64(30 * 24 * time.Hour)}

	var userID string
	userID, err = uss.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	// TODO: добавить отправку письма с подтверждением регистрации
	// TODO: добавить роль по умолчанию
	// TODO: вернуть токен

	user, err = uss.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}

	return &user_v1.SignUpResponse{Status: "success", Token: "", RoleName: ""}, nil
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

func (uss *userServiceServer) isUserExist(ctx context.Context, login string) bool {
	_, err := uss.repo.FindByLogin(ctx, login)
	return err == nil
}
