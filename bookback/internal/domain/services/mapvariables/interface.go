package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, event *entity.BookEvent) (string, error)
}
