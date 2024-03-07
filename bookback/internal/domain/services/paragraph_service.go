package services

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/converter"
)

type ParagraphService interface {
	CreateParagraph(ctx context.Context, request models.CreateParagraphRequest) (*models.Paragraph, error)
	GetParagraphByID(ctx context.Context, id string) (*models.Paragraph, error)
	UpdateParagraph(ctx context.Context, id string, request models.UpdateParagraphRequest) (*models.Paragraph, error)
	DeleteParagraph(ctx context.Context, id string) (*models.Paragraph, error)
	ListParagraphs(ctx context.Context, limit uint64, offset uint64) ([]*models.Paragraph, error)

	GetParagraphsByPageID(ctx context.Context, pageID string) ([]*models.Paragraph, error)
}

type paragraphService struct {
	repo SimpleRepo[*entity.Paragraph]
}

func NewParagraphService(repo SimpleRepo[*entity.Paragraph]) ParagraphService {
	return &paragraphService{repo}
}

func (s *paragraphService) CreateParagraph(ctx context.Context, request models.CreateParagraphRequest) (*models.Paragraph, error) {
	paragraph := converter.ParagraphModelToEntity(request.Paragraph)

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

	return converter.ParagraphEntityToModel(paragraph), nil
}

func (s *paragraphService) UpdateParagraph(ctx context.Context, id string, request models.UpdateParagraphRequest) (*models.Paragraph, error) {
	paragraph := converter.ParagraphModelToEntity(request.Paragraph)

	updatedParagraph, err := s.repo.Update(ctx, id, paragraph)
	if err != nil {
		return nil, err
	}

	return converter.ParagraphEntityToModel(updatedParagraph), nil
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

func (s *paragraphService) ListParagraphs(ctx context.Context, limit uint64, offset uint64) ([]*models.Paragraph, error) {
	paragraph, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return converter.ParagraphsEntityToModel(paragraph), nil
}

func (s *paragraphService) GetParagraphsByPageID(ctx context.Context, pageID string) ([]*models.Paragraph, error) {
	paragraphs, err := s.repo.FindByKV(ctx, "page_id", pageID)
	if err != nil {
		return nil, err
	}

	return converter.ParagraphsEntityToModel(paragraphs), nil
}
