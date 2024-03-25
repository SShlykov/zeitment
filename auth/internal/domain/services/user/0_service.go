package user

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

type Repository interface {
	List(ctx context.Context, options dbutils.Pagination) ([]*entity.User, error)
	Create(ctx context.Context, item *entity.User) (string, error)
	Update(ctx context.Context, id string, item *entity.User) (*entity.User, error)
	HardDelete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*entity.User, error)
	FindByLogin(ctx context.Context, login string) (*entity.User, error)
	Count(ctx context.Context) (uint64, error)
}

type Service struct {
	user_v1.UnimplementedUserServiceServer
	repo Repository
}

const (
	StatusError = "error"
)

func NewService(repository Repository) *Service {
	return &Service{repo: repository}
}

func NewUserServiceServer(repository Repository) user_v1.UserServiceServer {
	return &Service{repo: repository}
}

func (uss *Service) isUserExist(ctx context.Context, login string) bool {
	user, _ := uss.repo.FindByLogin(ctx, login)
	return user != nil
}

func (uss *Service) getPaginationMetadata(ctx context.Context, pagination *user_v1.Pagination) *user_v1.PaginationMetadata {
	page, pageSize := pagination.Page, pagination.PageSize
	count, _ := uss.repo.Count(ctx)
	totalPages := ((count - 1) / pageSize) + 1
	return &user_v1.PaginationMetadata{Page: page, PageSize: pageSize, Total: count, TotalPages: totalPages}
}
