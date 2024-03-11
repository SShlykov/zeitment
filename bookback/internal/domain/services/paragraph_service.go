package services

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/adapters"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/dbutils"
)

//go:generate mockgen -destination=../../../tests/mocks/domain/services/paragraph_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services ParagraphService
type ParagraphService interface {
	CreateParagraph(ctx context.Context, request models.CreateParagraphRequest) (*models.Paragraph, error)
	GetParagraphByID(ctx context.Context, id string) (*models.Paragraph, error)
	UpdateParagraph(ctx context.Context, id string, request models.UpdateParagraphRequest) (*models.Paragraph, error)
	DeleteParagraph(ctx context.Context, id string) (*models.Paragraph, error)
	ListParagraphs(ctx context.Context, request models.RequestParagraph) ([]*models.Paragraph, error)

	GetParagraphsByPageID(ctx context.Context, pageID string, request models.RequestParagraph) ([]*models.Paragraph, error)
}

type paragraphService struct {
	repo SimpleRepo[*entity.Paragraph]
}

func NewParagraphService(repo SimpleRepo[*entity.Paragraph]) ParagraphService {
	return &paragraphService{repo}
}

func (s *paragraphService) CreateParagraph(ctx context.Context, request models.CreateParagraphRequest) (*models.Paragraph, error) {
	paragraph := adapters.ParagraphModelToEntity(request.Paragraph)

	id, err := s.repo.Create(ctx, paragraph)
	if err != nil {
		return nil, err
	}

	return s.GetParagraphByID(ctx, id)
}

func (s *paragraphService) GetParagraphByID(ctx context.Context, id string) (*models.Paragraph, error) {
	paragraph, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return adapters.ParagraphEntityToModel(paragraph), nil
}

func (s *paragraphService) UpdateParagraph(ctx context.Context, id string,
	request models.UpdateParagraphRequest) (*models.Paragraph, error) {
	paragraph := adapters.ParagraphModelToEntity(request.Paragraph)

	updatedParagraph, err := s.repo.Update(ctx, id, paragraph)
	if err != nil {
		return nil, err
	}

	return adapters.ParagraphEntityToModel(updatedParagraph), nil
}

func (s *paragraphService) DeleteParagraph(ctx context.Context, id string) (*models.Paragraph, error) {
	paragraph, err := s.GetParagraphByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return paragraph, err
}

func (s *paragraphService) ListParagraphs(ctx context.Context, request models.RequestParagraph) ([]*models.Paragraph, error) {
	options := dbutils.NewPagination(&request.Options)

	paragraph, err := s.repo.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters.ParagraphsEntityToModel(paragraph), nil
}

func (s *paragraphService) GetParagraphsByPageID(ctx context.Context, pageID string,
	request models.RequestParagraph) ([]*models.Paragraph, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("page_id", pageID),
		dbutils.NewPagination(&request.Options),
	)

	paragraphs, err := s.repo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters.ParagraphsEntityToModel(paragraphs), nil
}
