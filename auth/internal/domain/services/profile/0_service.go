package token

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

type Repository interface {
	List(ctx context.Context, options dbutils.Pagination) ([]*entity.UserProfile, error)
	Create(ctx context.Context, item *entity.UserProfile) (string, error)
	Update(ctx context.Context, id string, item *entity.UserProfile) (*entity.UserProfile, error)
	HardDelete(ctx context.Context, id string) error
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*entity.UserProfile, error)
	FindByLogin(ctx context.Context, login string) (*entity.UserProfile, error)
	Count(ctx context.Context) (uint64, error)
}

type Service struct {
	repo Repository
}

func NewService(repository Repository) *Service {
	return &Service{repo: repository}
}
