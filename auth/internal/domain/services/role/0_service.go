package role

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

type Repository interface {
	List(ctx context.Context, options dbutils.Pagination) ([]*entity.Role, error)
	Create(ctx context.Context, item *entity.Role) (string, error)
	Update(ctx context.Context, id string, item *entity.Role) (*entity.Role, error)
	HardDelete(ctx context.Context, id string) error
	FindByKV(ctx context.Context, options dbutils.QueryOptions) ([]*entity.Role, error)
	FindByLogin(ctx context.Context, login string) (*entity.Role, error)
	Count(ctx context.Context) (uint64, error)
}

type Service struct {
	repo Repository
}

func NewService(repository Repository) *Service {
	return &Service{repo: repository}
}
