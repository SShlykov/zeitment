package page

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
)

type Service interface {
	CreatePage(ctx context.Context, page *models.Page) (*models.Page, error)
	GetPageByID(ctx context.Context, id string) (*models.Page, error)
	UpdatePage(ctx context.Context, id string, page *models.Page) (*models.Page, error)
	DeletePage(ctx context.Context, id string) (*models.Page, error)
	ListPages(ctx context.Context) ([]models.Page, error)

	GetPagesByChapterID(ctx context.Context, chapterID string) ([]models.Page, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreatePage(ctx context.Context, page *models.Page) (*models.Page, error) {
	id, err := s.repo.Create(ctx, page)
	if err != nil {
		return nil, err
	}
	page.ID = id

	return page, err
}

func (s *service) GetPageByID(ctx context.Context, id string) (*models.Page, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdatePage(ctx context.Context, id string, page *models.Page) (*models.Page, error) {
	return s.repo.Update(ctx, id, page)
}

func (s *service) DeletePage(ctx context.Context, id string) (*models.Page, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) ListPages(ctx context.Context) ([]models.Page, error) {
	return s.repo.List(ctx)
}

func (s *service) GetPagesByChapterID(ctx context.Context, chapterID string) ([]models.Page, error) {
	return s.repo.GetPagesByChapterID(ctx, chapterID)
}
