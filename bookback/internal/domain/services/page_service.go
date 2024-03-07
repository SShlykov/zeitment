package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/converter"
)

type PageService interface {
	CreatePage(ctx context.Context, request models.CreatePageRequest) (*models.Page, error)
	GetPageByID(ctx context.Context, id string) (*models.Page, error)
	UpdatePage(ctx context.Context, id string, prequest models.UpdatePageRequest) (*models.Page, error)
	DeletePage(ctx context.Context, id string) (*models.Page, error)
	ListPages(ctx context.Context, limit uint64, offset uint64) ([]*models.Page, error)

	GetPagesByChapterID(ctx context.Context, chapterID string) ([]*models.Page, error)
}

type pageService struct {
	repo SimpleRepo[*entity.Page]
}

func NewPageService(repo SimpleRepo[*entity.Page]) PageService {
	return &pageService{repo}
}

func (s *pageService) CreatePage(ctx context.Context, request models.CreatePageRequest) (*models.Page, error) {
	page := converter.PageModelToEntity(request.Page)

	id, err := s.repo.Create(ctx, page)
	if err != nil {
		return nil, err
	}

	return s.GetPageByID(ctx, id)
}

func (s *pageService) GetPageByID(ctx context.Context, id string) (*models.Page, error) {
	page, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.PageEntityToModel(page), nil
}

func (s *pageService) UpdatePage(ctx context.Context, id string, request models.UpdatePageRequest) (*models.Page, error) {
	page := converter.PageModelToEntity(request.Page)

	updatedPage, err := s.repo.Update(ctx, id, page)
	if err != nil {
		return nil, err
	}

	return converter.PageEntityToModel(updatedPage), nil
}

func (s *pageService) DeletePage(ctx context.Context, id string) (*models.Page, error) {
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

func (s *pageService) ListPages(ctx context.Context, limit uint64, offset uint64) ([]*models.Page, error) {
	pages, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return converter.PagesEntityToModel(pages), nil
}

func (s *pageService) GetPagesByChapterID(ctx context.Context, chapterID string) ([]*models.Page, error) {
	pages, err := s.repo.FindByKV(ctx, "chapter_id", chapterID)
	if err != nil {
		return nil, err
	}

	return converter.PagesEntityToModel(pages), nil
}
