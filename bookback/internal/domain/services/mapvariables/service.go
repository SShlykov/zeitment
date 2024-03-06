package mapvariables

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Service interface {
	CreateMapVariable(ctx context.Context, variable *entity.MapVariable) (*entity.MapVariable, error)
	GetMapVariableByID(ctx context.Context, id string) (*entity.MapVariable, error)
	UpdateMapVariable(ctx context.Context, id string, variable *entity.MapVariable) (*entity.MapVariable, error)
	DeleteMapVariable(ctx context.Context, id string) (*entity.MapVariable, error)

	GetMapVariablesByBookID(ctx context.Context, mapID string) ([]*entity.MapVariable, error)
	GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]*entity.MapVariable, error)
	GetMapVariablesByPageID(ctx context.Context, pageID string) ([]*entity.MapVariable, error)
	GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]*entity.MapVariable, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
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
	mapVariable, err := s.GetMapVariableByID(ctx, id)
	if err != nil {
		return nil, errors.Join(errors.New("MapVariable not found"), err)
	}

	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapVariable, err
}

func (s *service) GetMapVariablesByBookID(ctx context.Context, mapID string) ([]*entity.MapVariable, error) {
	return s.repo.FindByKV(ctx, "map_id", mapID)
}

func (s *service) GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]*entity.MapVariable, error) {
	return s.repo.FindByKV(ctx, "chapter_id", chapterID)
}

func (s *service) GetMapVariablesByPageID(ctx context.Context, pageID string) ([]*entity.MapVariable, error) {
	return s.repo.FindByKV(ctx, "page_id", pageID)
}

func (s *service) GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]*entity.MapVariable, error) {
	return s.repo.FindByKV(ctx, "paragraph_id", paragraphID)
}
