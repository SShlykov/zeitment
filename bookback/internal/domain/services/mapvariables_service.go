package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/adapters"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

//go:generate mockgen -destination=../../../tests/mocks/domain/services/mapvariables_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services MapVariablesService
type MapVariablesService interface {
	CreateMapVariable(ctx context.Context, request models.CreateMapVariableRequest) (*models.MapVariable, error)
	GetMapVariableByID(ctx context.Context, id string) (*models.MapVariable, error)
	UpdateMapVariable(ctx context.Context, id string, variable models.UpdateMapVariableRequest) (*models.MapVariable, error)
	DeleteMapVariable(ctx context.Context, id string) (*models.MapVariable, error)

	GetMapVariablesByBookID(ctx context.Context, mapID string, request models.RequestMapVariable) ([]*models.MapVariable, error)
	GetMapVariablesByChapterID(ctx context.Context, chapterID string, request models.RequestMapVariable) ([]*models.MapVariable, error)
	GetMapVariablesByPageID(ctx context.Context, pageID string, request models.RequestMapVariable) ([]*models.MapVariable, error)
	GetMapVariablesByParagraphID(ctx context.Context, paragraphID string, request models.RequestMapVariable) ([]*models.MapVariable, error)
}

type mapVariablesService struct {
	repo SimpleRepo[*entity.MapVariable]
}

func NewMapVariablesService(repo SimpleRepo[*entity.MapVariable]) MapVariablesService {
	return &mapVariablesService{repo}
}

func (s *mapVariablesService) CreateMapVariable(ctx context.Context, request models.CreateMapVariableRequest) (*models.MapVariable, error) {
	variable := adapters.MapVariableModelToEntity(request.MapVariable)

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

	return adapters.MapVariableEntityToModel(variable), nil
}

func (s *mapVariablesService) UpdateMapVariable(ctx context.Context, id string,
	request models.UpdateMapVariableRequest) (*models.MapVariable, error) {
	mapVariable := adapters.MapVariableModelToEntity(request.MapVariable)

	updatedVariable, err := s.repo.Update(ctx, id, mapVariable)
	if err != nil {
		return nil, err
	}

	return adapters.MapVariableEntityToModel(updatedVariable), nil
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

func (s *mapVariablesService) GetMapVariablesByBookID(ctx context.Context, mapID string,
	request models.RequestMapVariable) ([]*models.MapVariable, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("book_id", mapID),
		dbutils.NewPagination(&request.Options),
	)

	variable, err := s.repo.FindByKV(ctx, options)

	if err != nil {
		return nil, err
	}

	return adapters.MapVariablesEntityToModel(variable), nil
}

func (s *mapVariablesService) GetMapVariablesByChapterID(ctx context.Context, chapterID string,
	request models.RequestMapVariable) ([]*models.MapVariable, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("chapter_id", chapterID),
		dbutils.NewPagination(&request.Options),
	)

	variable, err := s.repo.FindByKV(ctx, options)

	if err != nil {
		return nil, err
	}

	return adapters.MapVariablesEntityToModel(variable), nil
}

func (s *mapVariablesService) GetMapVariablesByPageID(ctx context.Context, pageID string,
	request models.RequestMapVariable) ([]*models.MapVariable, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("page_id", pageID),
		dbutils.NewPagination(&request.Options),
	)

	variable, err := s.repo.FindByKV(ctx, options)

	if err != nil {
		return nil, err
	}

	return adapters.MapVariablesEntityToModel(variable), nil
}

func (s *mapVariablesService) GetMapVariablesByParagraphID(ctx context.Context, paragraphID string,
	request models.RequestMapVariable) ([]*models.MapVariable, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("paragraph_id", paragraphID),
		dbutils.NewPagination(&request.Options),
	)
	variable, err := s.repo.FindByKV(ctx, options)

	if err != nil {
		return nil, err
	}

	return adapters.MapVariablesEntityToModel(variable), nil
}
