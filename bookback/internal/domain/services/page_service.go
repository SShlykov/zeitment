package services

import (
	"context"
	"errors"
	"github.com/SShlykov/zeitment/bookback/internal/domain/adapters"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

//go:generate mockgen -destination=../../../tests/mocks/domain/services/page_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services PageService
type PageService interface {
	CreatePage(ctx context.Context, request models.CreatePageRequest) (*models.Page, error)
	GetPageByID(ctx context.Context, id string) (*models.Page, error)
	UpdatePage(ctx context.Context, id string, request models.UpdatePageRequest) (*models.Page, error)
	DeletePage(ctx context.Context, id string) (*models.Page, error)
	ListPages(ctx context.Context, request models.RequestPage) ([]*models.Page, error)

	GetPagesByChapterID(ctx context.Context, chapterID string, request models.RequestPage) ([]*models.Page, error)

	TogglePublic(ctx context.Context, request models.TogglePageRequest) (*models.Page, error)
}

type pageService struct {
	repo SimpleRepo[*entity.Page]
}

func NewPageService(repo SimpleRepo[*entity.Page]) PageService {
	return &pageService{repo}
}

func (s *pageService) TogglePublic(ctx context.Context, request models.TogglePageRequest) (*models.Page, error) {
	page, err := s.repo.FindByID(ctx, request.PageID)
	if err != nil {
		return nil, err
	}
	page.IsPublic = !page.IsPublic

	var updated *entity.Page
	updated, err = s.repo.Update(ctx, request.PageID, page)
	if err != nil {
		return nil, err
	}

	return adapters.PageEntityToModel(updated), nil
}

func (s *pageService) CreatePage(ctx context.Context, request models.CreatePageRequest) (*models.Page, error) {
	page := adapters.PageModelToEntity(request.Page)

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

	return adapters.PageEntityToModel(page), nil
}

func (s *pageService) UpdatePage(ctx context.Context, id string, request models.UpdatePageRequest) (*models.Page, error) {
	page := adapters.PageModelToEntity(request.Page)

	updatedPage, err := s.repo.Update(ctx, id, page)
	if err != nil {
		return nil, err
	}

	return adapters.PageEntityToModel(updatedPage), nil
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

func (s *pageService) ListPages(ctx context.Context, request models.RequestPage) ([]*models.Page, error) {
	options := dbutils.NewPagination(&request.Options)

	pages, err := s.repo.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters.PagesEntityToModel(pages), nil
}

func (s *pageService) GetPagesByChapterID(ctx context.Context, chapterID string, request models.RequestPage) ([]*models.Page, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("chapter_id", chapterID),
		dbutils.NewPagination(&request.Options),
	)

	pages, err := s.repo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters.PagesEntityToModel(pages), nil
}
