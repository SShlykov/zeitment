package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/converter"
)

type MapVariablesService interface {
	CreateMapVariable(ctx context.Context, request models.CreateMapVariableRequest) (*models.MapVariable, error)
	GetMapVariableByID(ctx context.Context, id string) (*models.MapVariable, error)
	UpdateMapVariable(ctx context.Context, id string, variable models.UpdateMapVariableRequest) (*models.MapVariable, error)
	DeleteMapVariable(ctx context.Context, id string) (*models.MapVariable, error)

	GetMapVariablesByBookID(ctx context.Context, mapID string) ([]*models.MapVariable, error)
	GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]*models.MapVariable, error)
	GetMapVariablesByPageID(ctx context.Context, pageID string) ([]*models.MapVariable, error)
	GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]*models.MapVariable, error)
}

type mapVariablesService struct {
	repo SimpleRepo[*entity.MapVariable]
}

func NewMapVariablesService(repo SimpleRepo[*entity.MapVariable]) MapVariablesService {
	return &mapVariablesService{repo}
}

func (s *mapVariablesService) CreateMapVariable(ctx context.Context, request models.CreateMapVariableRequest) (*models.MapVariable, error) {
	variable := converter.MapVariableModelToEntity(request.MapVariable)

	id, err := s.repo.Create(ctx, variable)
	if err != nil {
		return nil, err
	}

	return s.GetMapVariableByID(ctx, id)
}

func (s *mapVariablesService) GetMapVariableByID(ctx context.Context, id string) (*models.MapVariable, error) {
	variable, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.MapVariableEntityToModel(variable), nil
}

func (s *mapVariablesService) UpdateMapVariable(ctx context.Context, id string, request models.UpdateMapVariableRequest) (*models.MapVariable, error) {
	mapVariable := converter.MapVariableModelToEntity(request.MapVariable)

	updatedVariable, err := s.repo.Update(ctx, id, mapVariable)
	if err != nil {
		return nil, err
	}

	return converter.MapVariableEntityToModel(updatedVariable), nil
}

func (s *mapVariablesService) DeleteMapVariable(ctx context.Context, id string) (*models.MapVariable, error) {
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

func (s *mapVariablesService) GetMapVariablesByBookID(ctx context.Context, mapID string) ([]*models.MapVariable, error) {
	variable, err := s.repo.FindByKV(ctx, "map_id", mapID)

	if err != nil {
		return nil, err
	}

	return converter.MapVariablesEntityToModel(variable), nil
}

func (s *mapVariablesService) GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]*models.MapVariable, error) {
	variable, err := s.repo.FindByKV(ctx, "chapter_id", chapterID)

	if err != nil {
		return nil, err
	}

	return converter.MapVariablesEntityToModel(variable), nil
}

func (s *mapVariablesService) GetMapVariablesByPageID(ctx context.Context, pageID string) ([]*models.MapVariable, error) {
	variable, err := s.repo.FindByKV(ctx, "page_id", pageID)

	if err != nil {
		return nil, err
	}

	return converter.MapVariablesEntityToModel(variable), nil
}

func (s *mapVariablesService) GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]*models.MapVariable, error) {
	variable, err := s.repo.FindByKV(ctx, "paragraph_id", paragraphID)

	if err != nil {
		return nil, err
	}

	return converter.MapVariablesEntityToModel(variable), nil
}
