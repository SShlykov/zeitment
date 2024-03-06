package paragraph

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Service interface {
	CreateParagraph(ctx context.Context, paragraph *entity.Paragraph) (*entity.Paragraph, error)
	GetParagraphByID(ctx context.Context, id string) (*entity.Paragraph, error)
	UpdateParagraph(ctx context.Context, id string, paragraph *entity.Paragraph) (*entity.Paragraph, error)
	DeleteParagraph(ctx context.Context, id string) (*entity.Paragraph, error)
	ListParagraphs(ctx context.Context, limit uint64, offset uint64) ([]*entity.Paragraph, error)

	GetParagraphsByPageID(ctx context.Context, pageID string) ([]*entity.Paragraph, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateParagraph(ctx context.Context, paragraph *entity.Paragraph) (*entity.Paragraph, error) {
	id, err := s.repo.Create(ctx, paragraph)
	if err != nil {
		return nil, err
	}
	paragraph.ID = id

	return paragraph, err
}

func (s *service) GetParagraphByID(ctx context.Context, id string) (*entity.Paragraph, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateParagraph(ctx context.Context, id string, paragraph *entity.Paragraph) (*entity.Paragraph, error) {
	return s.repo.Update(ctx, id, paragraph)
}

func (s *service) DeleteParagraph(ctx context.Context, id string) (*entity.Paragraph, error) {
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

func (s *service) ListParagraphs(ctx context.Context, limit uint64, offset uint64) ([]*entity.Paragraph, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *service) GetParagraphsByPageID(ctx context.Context, pageID string) ([]*entity.Paragraph, error) {
	return s.repo.FindByKV(ctx, "page_id", pageID)
}
