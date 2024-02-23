package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
)

type Service interface {
	CreateParagraph(ctx context.Context, paragraph *models.Paragraph) (*models.Paragraph, error)
	GetParagraphByID(ctx context.Context, id string) (*models.Paragraph, error)
	UpdateParagraph(ctx context.Context, id string, paragraph *models.Paragraph) (*models.Paragraph, error)
	DeleteParagraph(ctx context.Context, id string) (*models.Paragraph, error)
	ListParagraphs(ctx context.Context, pageID string) ([]models.Paragraph, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateParagraph(ctx context.Context, paragraph *models.Paragraph) (*models.Paragraph, error) {
	id, err := s.repo.Create(ctx, paragraph)
	if err != nil {
		return nil, err
	}
	paragraph.ID = id

	return paragraph, err
}

func (s *service) GetParagraphByID(ctx context.Context, id string) (*models.Paragraph, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateParagraph(ctx context.Context, id string, paragraph *models.Paragraph) (*models.Paragraph, error) {
	return s.repo.Update(ctx, id, paragraph)
}

func (s *service) DeleteParagraph(ctx context.Context, id string) (*models.Paragraph, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) ListParagraphs(ctx context.Context, _ string) ([]models.Paragraph, error) {
	return s.repo.List(ctx)
}
