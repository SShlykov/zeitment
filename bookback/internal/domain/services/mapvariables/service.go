package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/adapters/db/postgres/mapvarrepo"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Service interface {
	CreateMapVariable(ctx context.Context, variable *entity.MapVariable) (*entity.MapVariable, error)
	GetMapVariableByID(ctx context.Context, id string) (*entity.MapVariable, error)
	UpdateMapVariable(ctx context.Context, id string, variable *entity.MapVariable) (*entity.MapVariable, error)
	DeleteMapVariable(ctx context.Context, id string) (*entity.MapVariable, error)
	GetMapVariablesByBookID(ctx context.Context, mapID string) ([]entity.MapVariable, error)
	GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]entity.MapVariable, error)
	GetMapVariablesByPageID(ctx context.Context, pageID string) ([]entity.MapVariable, error)
	GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]entity.MapVariable, error)
}

type service struct {
	repo mapvarrepo.Repository
}

func NewService(repo mapvarrepo.Repository) Service {
	return &service{repo}
}

func (s *service) CreateMapVariable(ctx context.Context, variable *entity.MapVariable) (*entity.MapVariable, error) {
	id, err := s.repo.Create(ctx, variable)
	if err != nil {
		return nil, err
	}
	variable.ID = id

	return variable, err
}

func (s *service) GetMapVariableByID(ctx context.Context, id string) (*entity.MapVariable, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateMapVariable(ctx context.Context, id string, variable *entity.MapVariable) (*entity.MapVariable, error) {
	return s.repo.Update(ctx, id, variable)
}

func (s *service) DeleteMapVariable(ctx context.Context, id string) (*entity.MapVariable, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetMapVariablesByBookID(ctx context.Context, mapID string) ([]entity.MapVariable, error) {
	return s.repo.GetByBookID(ctx, mapID)
}

func (s *service) GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]entity.MapVariable, error) {
	return s.repo.GetByChapterID(ctx, chapterID)
}

func (s *service) GetMapVariablesByPageID(ctx context.Context, pageID string) ([]entity.MapVariable, error) {
	return s.repo.GetByPageID(ctx, pageID)
}

func (s *service) GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]entity.MapVariable, error) {
	return s.repo.GetByParagraphID(ctx, paragraphID)
}
