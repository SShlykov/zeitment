package page

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Service interface {
	CreatePage(ctx context.Context, page *entity.Page) (*entity.Page, error)
	GetPageByID(ctx context.Context, id string) (*entity.Page, error)
	UpdatePage(ctx context.Context, id string, page *entity.Page) (*entity.Page, error)
	DeletePage(ctx context.Context, id string) (*entity.Page, error)
	ListPages(ctx context.Context, limit uint64, offset uint64) ([]*entity.Page, error)

	GetPagesByChapterID(ctx context.Context, chapterID string) ([]*entity.Page, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreatePage(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	id, err := s.repo.Create(ctx, page)
	if err != nil {
		return nil, err
	}
	page.ID = id

	return page, err
}

func (s *service) GetPageByID(ctx context.Context, id string) (*entity.Page, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdatePage(ctx context.Context, id string, page *entity.Page) (*entity.Page, error) {
	return s.repo.Update(ctx, id, page)
}

func (s *service) DeletePage(ctx context.Context, id string) (*entity.Page, error) {
	page, err := s.GetPageByID(ctx, id)
	if err != nil {
		return nil, errors.Join(errors.New("page not found"), err)
	}
	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return page, err
}

func (s *service) ListPages(ctx context.Context, limit uint64, offset uint64) ([]*entity.Page, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *service) GetPagesByChapterID(ctx context.Context, chapterID string) ([]*entity.Page, error) {
	return s.repo.FindByKV(ctx, "chapter_id", chapterID)
}
